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
func (repo *Repo) GetAllTasks() engine.Tasks {
	// Get all todos
	toDos, err := repo.Db.(*sql.DB).Query("SELECT * FROM todotasks ;")

	if err != nil {
		fmt.Println(err)
	}

	// Get all inprogress
	inprogress, err := repo.Db.(*sql.DB).Query("SELECT * FROM inprogresstasks ;")

	if err != nil {
		fmt.Println(err)
	}

	// Get all done
	done, err := repo.Db.(*sql.DB).Query("SELECT * FROM donetasks ;")

	if err != nil {
		fmt.Println(err)
	}

	// Build tasks
	tasks := engine.Tasks{
		ToDo:       engine.QueryToTickets(toDos),
		InProgress: engine.QueryToTickets(inprogress),
		Done:       engine.QueryToTickets(done),
	}

	return tasks
}

// Open creates the database connection used for the kanban board
func Open() *sql.DB {
	db, err := sql.Open("sqlite3", "../server/database/repository/repo.db")

	if err != nil {
		fmt.Println(err)
	}

	return db
}
