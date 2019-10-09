package repository

import (
	"database/sql"
	"fmt"
	kb "kanbanBoard"
	"os"
	"path"

	"github.com/google/uuid"

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

	q, err := db.Exec(initState)
	fmt.Println(err)
	fmt.Println(q)
	r.db = db

	return nil
}

// GetAllTasks pulls all tasks from the database and converts them
// into Tasks struct
func (r Repo) GetAllTasks() kb.Tasks {

	// Get all todos
	allTodos, _ := r.db.Query("SELECT title, desc, state, deadline, id FROM tasks")
	tasks := kb.Tasks{}
	var title, desc, state, deadline, id string

	for allTodos.Next() {
		err := allTodos.Scan(&title, &desc, &state, &deadline, &id)
		if err != nil {
			fmt.Println(err)
		}

		// Distinguish between states
		switch state {
		case "todo":
			tasks.ToDo = append(tasks.ToDo, kb.Ticket{
				Title:       title,
				Description: desc,
				Deadline:    deadline,
				ID:          id,
			})
			break

		case "inprogress":
			tasks.InProgress = append(tasks.InProgress, kb.Ticket{
				Title:       title,
				Description: desc,
				Deadline:    deadline,
				ID:          id,
			})
			break

		case "done":
			tasks.Done = append(tasks.Done, kb.Ticket{
				Title:       title,
				Description: desc,
				Deadline:    deadline,
				ID:          id,
			})
			break
		}
	}

	return tasks
}

// UpdateTicketState changes the state of a given ticket
func (r Repo) UpdateTicketState(state, id string) error {

	query, _ := r.db.Prepare(`UPDATE tasks
								SET state = ? 
								WHERE id = ? ;`)

	res, err := query.Exec(state, id)
	n, _ := res.RowsAffected()
	fmt.Println("Updated", n, "rows")

	return err
}

// SetTicketAsDoneAndDelete sets given ticket as done and removes it
// from the database
func (r Repo) SetTicketAsDoneAndDelete(id string) error {

	query, _ := r.db.Prepare(`DELETE FROM tasks
								WHERE state = 'done' 
								AND id = ? ;`)

	res, err := query.Exec(id)
	n, _ := res.RowsAffected()
	fmt.Println("Updated", n, "rows")

	return err
}

// AddNewTicket puts a ticket with given title and desc
// into the database
func (r Repo) AddNewTicket(title, desc, deadline string) error {

	// Transfer data to inprogress
	query, err := r.db.Prepare(`INSERT INTO tasks
										VALUES (
												?, 
												?, 
												?, 
												?,
												?
											) ;`)

	fmt.Println(err)

	id := uuid.New().String()
	res, err := query.Exec(title, desc, "todo", deadline, id)
	n, _ := res.RowsAffected()
	fmt.Println("title:", title, "desc:", desc, "id:", id)
	fmt.Println("Updated", n, "rows")

	return err
}

// ClearDatabase deletes all entries with empty data
func (r Repo) ClearDatabase() error {
	query, _ := r.db.Prepare(`DELETE FROM tasks WHERE 
							state = '' OR
							desc = '' OR
							deadline = '' OR
							id = '' 
							;`)

	res, err := query.Exec()
	n, _ := res.RowsAffected()
	fmt.Println("Updated", n, "rows")

	return err
}

// Query to setup the database
const initState = `CREATE TABLE IF NOT EXISTS 'tasks' (
						'title' 	VARCHAR(64),
						'desc'  	VARCHAR(256), 
						'state' 	VARCHAR(64),
						'deadline' 	VARCHAR(16),
						'id'    	VARCHAR(16) 
						) ;`
