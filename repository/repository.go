package repository

import (
	"database/sql"
	"fmt"
	kb "kanbanBoard"
	"os"
	"strconv"

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
	r := &Repo{}

	return r, r.init(path)
}

func (r *Repo) init(path string) error {

	// Check if file does not exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("Database does not exist")
		fmt.Println("Creating a new database...")

		// Create database file
		os.Create(path)

		// Open databse
		db, err := sql.Open("sqlite3", path)

		fmt.Println(err)

		// Create table
		db.Exec(initState)

		// Define database in r
		r.db = db
	} else {
		db, err := sql.Open("sqlite3", path)

		fmt.Println(err)

		r.db = db
	}

	return nil
}

// GetAllTasks pulls all tasks from the database and converts them
// into Tasks struct
func (r Repo) GetAllTasks() kb.Tasks {

	// Get all todos
	allTodos, _ := r.db.Query("SELECT * FROM tasks")

	r.db.Close()

	tasks := kb.Tasks{}

	var title, desc, state, id string

	for allTodos.Next() {

		err := allTodos.Scan(&title, &desc, &state, &id)

		if err != nil {
			fmt.Println(err)
		}

		// Distinguish between states
		switch state {
		case "todo":
			tasks.ToDo = append(tasks.ToDo, kb.Ticket{Title: title, Description: desc, ID: id})
			break
		case "inprogress":
			tasks.InProgress = append(tasks.InProgress, kb.Ticket{Title: title, Description: desc, ID: id})
			break
		case "done":
			tasks.Done = append(tasks.Done, kb.Ticket{Title: title, Description: desc, ID: id})
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

	r.db.Close()

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

	r.db.Close()

	return err
}

// AddNewTicket puts a ticket with given title and desc
// into the database
func (r Repo) AddNewTicket(title, desc string) error {
	// Transfer data to inprogress
	query, err := r.db.Prepare(`INSERT INTO tasks
										VALUES (
												?, 
												?, 
												?, 
												?
											) ;`)

	fmt.Println(err)

	id := generateID()

	res, err := query.Exec(title, desc, "todo", id)
	n, _ := res.RowsAffected()

	fmt.Println("title:", title, "desc:", desc, "id:", id)

	fmt.Println("Updated", n, "rows")

	r.db.Close()

	return err
}

// ID = time in milliseconds
func generateID() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

const initState = `CREATE TABLE 'tasks' (
						'title' VARCHAR(64),
						'desc'  VARCHAR(256), 
						'state' VARCHAR(64),
						'id'    VARCHAR(64) 
						) ;`
