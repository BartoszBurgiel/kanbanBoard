package server

import (
	"database/sql"
	"kanbanBoard/server/engine"
)

// SqliteRepository handles a sqlite database
type SqliteRepository interface {
	GetTable(table string) *sql.Rows
	GetAllTasks() engine.Tasks
	TransferToInProgress() error
	TransferToDone() error
	SetAsDone() error
	AddTicket() error
}
