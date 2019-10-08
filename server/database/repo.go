package database

import (
	"database/sql"
	"fmt"
	"kanbanBoard/server/engine"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Repo represents a data base with its all methods
type Repo struct {
	Db interface{}
}

// NewRepo creates a new repository
func NewRepo(path string) *Repo {
	db, err := sql.Open("sqlite3", path)

	if err != nil {
		panic(err)
	}

	return &Repo{db}
}

// GetAllTasks pulls all tasks from the database and converts them
// into Tasks struct
func (r Repo) GetAllTasks() engine.Tasks {

	tasks := engine.Tasks{}

	// Get all todos
	allTodos, _ := r.Db.(*sql.DB).Query("SELECT * FROM tasks")

	var title, desc, state, id string

	for allTodos.Next() {

		err := allTodos.Scan(&title, &desc, &state, &id)

		if err != nil {
			fmt.Println(err)
		}

		// Distinguish between states
		switch state {
		case "todo":
			tasks.ToDo = append(tasks.ToDo, engine.Ticket{Title: title, Description: desc, ID: id})
			break
		case "inprogress":
			tasks.InProgress = append(tasks.InProgress, engine.Ticket{Title: title, Description: desc, ID: id})
			break
		case "done":
			tasks.Done = append(tasks.Done, engine.Ticket{Title: title, Description: desc, ID: id})
			break
		}
	}

	return tasks
}

// ChangeState changes the state of a given ticket
func (r Repo) ChangeState(state, id string) error {

	query, _ := r.Db.(*sql.DB).Prepare(`UPDATE tasks
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

	query, _ := r.Db.(*sql.DB).Prepare(`DELETE FROM tasks
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
	query, _ := r.Db.(*sql.DB).Prepare(`INSERT INTO tasks
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
