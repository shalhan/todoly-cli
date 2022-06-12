package main

import (
    "log"

	"todoly.app/todoly/board/repo/sqlite"
	taskSqlite "todoly.app/todoly/task/repo/sqlite"
	"todoly.app/todoly/cmd"

    _ "github.com/mattn/go-sqlite3"
)

const fileName = "sqlite.db"

func main() {
    boardRepo := sqlite.NewBoardRepo()
	taskRepo := taskSqlite.NewTaskRepo()

	if err := boardRepo.Migrate(); err != nil {
        log.Fatal(err)
    }

	if err := taskRepo.Migrate(); err != nil {
        log.Fatal(err)
    }

    
	cmd.Execute()	
}
