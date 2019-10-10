package repository

import (
	"database/sql"
	"fmt"
	"os"
	"path"

	// Sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

// Repo represents a data base with its all methods
type Repo struct {
	db   *sql.DB
	path string
}

// New creates a new repository
func New(path string) (*Repo, error) {
	r := &Repo{
		path: path,
	}

	return r, r.init()
}

func (r *Repo) init() error {

	p := path.Dir(r.path)
	if err := os.MkdirAll(p, os.ModePerm); err != nil {
		return err
	}

	db, err := sql.Open("sqlite3", r.path)
	if err != nil {
		return err
	}

	q, err := db.Exec(initColumns)
	fmt.Println(err)
	fmt.Println(q)

	q, err = db.Exec(newTickets)
	fmt.Println(err)
	fmt.Println(q)

	q, err = db.Exec(basicColumns)
	fmt.Println(err)
	fmt.Println(q)

	r.db = db

	return nil
}

// ClearDatabase deletes all entries with empty data
func (r Repo) ClearDatabase() error {
	query, _ := r.db.Prepare(`DELETE FROM tasks WHERE 
							state = '' OR
							desc = '' OR
							deadline = '' OR
							priority = '' OR
							id = '' 
							;`)

	res, err := query.Exec()
	n, _ := res.RowsAffected()
	fmt.Println("Updated", n, "rows")

	return err
}
