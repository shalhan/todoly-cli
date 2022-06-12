package sqlite

import (
	"database/sql"
	"errors"
	"log"

	"todoly.app/todoly/board/model"

	"github.com/mattn/go-sqlite3"
)

var (
    ErrDuplicate    = errors.New("record already exists")
    ErrNotExists    = errors.New("row not exists")
    ErrUpdateFailed = errors.New("update failed")
    ErrDeleteFailed = errors.New("delete failed")
)

type SQLiteRepository struct {
	db *sql.DB
}


func NewBoardRepo() *SQLiteRepository {
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		log.Fatal(err)
	}

	return &SQLiteRepository{
		db: db,
	}

}

func (r *SQLiteRepository) Migrate() error {
    query := `
    CREATE TABLE IF NOT EXISTS boards(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name VARCHAR NOT NULL UNIQUE,
		is_active BOOLEAN DEFAULT false,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP NULLABLE DEFAULT NULL
    );
    `

    _, err := r.db.Exec(query)
    return err
}

func (r *SQLiteRepository) Rollback() error {
    query := `DROP TABLE boards`

    _, err := r.db.Exec(query)
    return err
}


func (r *SQLiteRepository) Create(board model.Board) (*model.Board, error) {
    res, err := r.db.Exec("INSERT INTO boards(name) values(?)", board.Name)
    if err != nil {
        var sqliteErr sqlite3.Error
        if errors.As(err, &sqliteErr) {
            if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
                return nil, ErrDuplicate
            }
        }
        return nil, err
    }

    id, err := res.LastInsertId()
    if err != nil {
        return nil, err
    }

	board.ID = id

    return &board, nil
}

func (r *SQLiteRepository) All() ([]model.Board, error) {
    rows, err := r.db.Query("SELECT * FROM boards where deleted_at is null")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var all []model.Board

    for rows.Next() {
        var board model.Board 
        if err := rows.Scan(&board.ID, &board.Name, &board.IsActive, &board.CreatedAt, &board.DeletedAt); err != nil {
            return nil, err
		}
        all = append(all, board)
	}

    return all, nil
}

func (r *SQLiteRepository) Delete(id string) error {
    res, err := r.db.Exec("UPDATE boards SET deleted_at = CURRENT_TIMESTAMP where name = ?", id)
    if err != nil {
        return err
    }

    rowsAffected, err := res.RowsAffected()
    if err != nil {
        return err
    }

    if rowsAffected == 0 {
        return ErrDeleteFailed
    }

    return err
}

func (r *SQLiteRepository) Active(id string) error {
	_, err := r.db.Exec("UPDATE boards SET is_active = false where is_active = true")
	if err != nil {
		return err
	}
	_, err2 := r.db.Exec("UPDATE boards SET is_active = true where name  = ?", id)
	if err2 != nil {
		return err
	}

	return err
}

func (r *SQLiteRepository) GetActive() (*model.Board, error) {
    row := r.db.QueryRow("SELECT * FROM boards WHERE is_active = true AND deleted_at is null")

    var board model.Board
    if err := row.Scan(&board.ID, &board.Name, &board.IsActive, &board.CreatedAt, &board.DeletedAt); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, ErrNotExists
        }
        return nil, err
    }
    return &board, nil
}
