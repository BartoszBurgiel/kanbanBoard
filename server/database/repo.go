package database

import (
	"database/sql"
	"fmt"
	"kanbanBoard/server/engine"

	_ "github.com/mattn/go-sqlite3"
)

// Repo represents a data base with its all methods
type Repo struct {
	Db interface{}
}

func (r Repo) GetTable(table string) *sql.Rows {
	rows, err := r.Db.(*sql.DB).Query(fmt.Sprintf("SELECT * FROM %s ;", table))

	if err != nil {
		fmt.Println(err)
	}

	return rows
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
	// Get all todos
	toDos := r.GetTable("todotasks")

	// Get all inprogress
	inprogress := r.GetTable("inprogresstasks")

	// Get all done
	done := r.GetTable("donetasks")

	// Build tasks
	tasks := engine.Tasks{
		ToDo:       engine.QueryToTickets(toDos),
		InProgress: engine.QueryToTickets(inprogress),
		Done:       engine.QueryToTickets(done),
	}

	return tasks
}

func (r Repo) TransferToInProgress(id string) error {

	// Transfer data to inprogress

	r.Db.(*sql.DB).Query(`INSERT INTO 'inprogresstasks
							VALUES (`)

	return nil
}
