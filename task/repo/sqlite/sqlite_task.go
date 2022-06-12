package sqlite

import (
	"database/sql"
	"errors"
	"log"
	"strings"

	"todoly.app/todoly/task/model"

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


func NewTaskRepo() *SQLiteRepository {
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
    CREATE TABLE IF NOT EXISTS tasks(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title VARCHAR NOT NULL,
		brief TEXT NULLABLE,
		board_id INTEGER NOT NULL,
		status VARCHAR DEFAULT "Backlog",
		is_active BOOLEAN DEFAULT false,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP NULLABLE DEFAULT 'unixepoch',
		FOREIGN KEY(board_id) REFERENCES boards(id)
    );
    `

    _, err := r.db.Exec(query)
    return err
}

func (r *SQLiteRepository) Rollback() error {
    query := `DROP TABLE tasks`

    _, err := r.db.Exec(query)
    return err
}


func (r *SQLiteRepository) Create(task model.Task) (*model.Task, error) {
    res, err := r.db.Exec("INSERT INTO tasks(title, brief, board_id) values(?,?,?)", task.Title, task.Brief, task.BoardId)
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

	task.ID = id

    return &task, nil
}

func (r *SQLiteRepository) All(boardId int64) ([]model.Task, error) {
    rows, err := r.db.Query("SELECT * FROM tasks where board_id = ? and deleted_at = 'unixepoch'", boardId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var all []model.Task

    for rows.Next() {
        var task model.Task 
        if err := rows.Scan(
			&task.ID, 
			&task.Title, 
			&task.Brief, 
			&task.BoardId,
			&task.Status,
			&task.IsActive, 
			&task.CreatedAt, 
			&task.DeletedAt,
		); err != nil {
            return nil, err
		}
        all = append(all, task)
	}

    return all, nil
}

func (r *SQLiteRepository) Delete(ids []string) error {
	idVal := "(" + strings.Join(ids,",") + ")"

    res, err := r.db.Exec("UPDATE tasks SET deleted_at = CURRENT_TIMESTAMP where id in " + idVal)
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
	_, err := r.db.Exec("UPDATE tasks SET is_active = false where is_active = true")
	if err != nil {
		return err
	}
	_, err2 := r.db.Exec("UPDATE tasks SET is_active = true where name  = ?", id)
	if err2 != nil {
		return err
	}

	return err
}

func (r *SQLiteRepository) GetActive() (*model.Task, error) {
    row := r.db.QueryRow("SELECT * FROM tasks WHERE is_activate = true AND deleted_at is null")

    var task model.Task

    if err := row.Scan(
			&task.ID, 
			&task.Title, 
			&task.Brief, 
			&task.BoardId,
			&task.Status,
			&task.IsActive, 
			&task.CreatedAt, 
			&task.DeletedAt,
		);  err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, ErrNotExists
        }
        return nil, err
    }
    return &task, nil
}

func (r *SQLiteRepository) SetDone(ids []string) error {
	idVal := "(" + strings.Join(ids,",") + ")"
	_, err := r.db.Exec("UPDATE tasks set STATUS = 'Done' WHERE id IN " + idVal +  " and deleted_at = 'unixepoch'")
	
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (r *SQLiteRepository) SetAbort(ids []string) error {
	idVal := "(" + strings.Join(ids,",") + ")"
	_, err := r.db.Exec("UPDATE tasks set STATUS = 'Abort' WHERE id IN " + idVal +  " and deleted_at = 'unixepoch'")
	
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (r *SQLiteRepository) SetGoing(ids []string) error {
	idVal := "(" + strings.Join(ids,",") + ")"
	_, err := r.db.Exec("UPDATE tasks set STATUS = 'Going' WHERE id IN " + idVal +  " and deleted_at = 'unixepoch'")
	
	if err != nil {
		log.Fatal(err)
	}

	return err
}
