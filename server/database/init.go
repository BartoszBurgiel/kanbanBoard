package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Open creates the database connection used for the kanban board
func Open() *sql.DB {
	db, err := sql.Open("sqlite3", "../server/database/repository/repo.db")

	if err != nil {
		fmt.Println(err)
	}

	return db
}
