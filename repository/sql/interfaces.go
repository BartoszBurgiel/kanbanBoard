package sql

import "database/sql"

// SqliteRepository handles a sqlite database
type SqliteRepository interface {
	GetAllTasks() *sql.Rows
	ChangeState(state, id string) error
	SetAsDone(id string) error
	AddTicket(title, desc string) error
	ChangeTicket(state, id string)
}
