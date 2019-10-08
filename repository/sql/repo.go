package sql

import (
	"database/sql"
	"fmt"
	"time"

	// Sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

// Repo represents a data base with its all methods
type Repo struct {
	db *sql.DB
}

// NewRepo creates a new repository
func NewRepo(path string) (*Repo, error) {
	db, err := sql.Open("sqlite3", path)

	if err != nil {
		panic(err)
	}

	r := &Repo{db}

	return r, r.init()
}

func (r *Repo) init() error {
	//init
	//db exists?
	//no-> create db init tables

	//db.Exec(initstate)

	return nil
}

// GetAllTasks pulls all tasks from the database and converts them
// into Tasks struct
func (r Repo) GetAllTasks() *sql.Rows {

	// Get all todos
	allTodos, _ := r.db.Query("SELECT * FROM tasks")

	return allTodos
}

// ChangeState changes the state of a given ticket
func (r Repo) ChangeState(state, id string) error {

	query, _ := r.db.Prepare(`UPDATE tasks
								SET state = ? 
								WHERE id = ? ;`)

	res, err := query.Exec(state, id)
	n, _ := res.RowsAffected()

	fmt.Println("Updated", n, "rows")

	return err
}

// SetAsDone sets given ticket as done and removes it
// from the database
func (r Repo) SetAsDone(id string) error {

	query, _ := r.db.Prepare(`DELETE FROM tasks
										WHERE state = 'done' 
										AND id = ? ;`)

	res, err := query.Exec(id)
	n, _ := res.RowsAffected()

	fmt.Println("Updated", n, "rows")

	return err
}

// AddTicket puts a ticket with given title and desc
// into the database
func (r Repo) AddTicket(title, desc string) error {
	// Transfer data to inprogress
	query, _ := r.db.Prepare(`INSERT INTO tasks
										VALUES (
												?, 
												?, 
												?, 
												?
											) ;`)

	res, err := query.Exec(title, desc, "todo", generateID())
	n, _ := res.RowsAffected()

	fmt.Println("Updated", n, "rows")

	return err
}

// ID = time in milliseconds
func generateID() string {
	return string(time.Now().UnixNano())
}

const initstate = `Create table xy if not exist`
