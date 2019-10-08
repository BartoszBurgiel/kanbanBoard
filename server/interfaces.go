package server

import (
	"database/sql"
	"kanbanBoard/server/engine"
)

// SqliteRepository handles a sqlite database
type SqliteRepository interface {
	GetTable(table string) *sql.Rows
	GetAllTasks() engine.Tasks
	TransferToInProgress(id string) error
	TransferToDone(id string) error
	SetAsDone(id string) error
	AddTicket() error
}
