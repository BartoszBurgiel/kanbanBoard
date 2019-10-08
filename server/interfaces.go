package server

import (
	"kanbanBoard/server/engine"
)

// SqliteRepository handles a sqlite database
type SqliteRepository interface {
	GetAllTasks() engine.Tasks
	ChangeState(state, id string) error
	SetAsDone(id string) error
	AddTicket() error
}
