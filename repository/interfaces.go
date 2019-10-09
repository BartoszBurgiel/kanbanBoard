package repository

import (
	kb "kanbanBoard"
	"time"
)

// SqliteRepository handles a sqlite database
type SqliteRepository interface {
	GetAllTasks() kb.Tasks
	UpdateTicketState(newState, id string) error
	SetTicketAsDoneAndDelete(id string) error
	AddNewTicket(title, desc string, deadline time.Time, priority string) error
	HandleTicketEvent(state, id string) error
	ClearDatabase() error
}
