package repository

import kb "kanbanBoard"

// SqliteRepository handles a sqlite database
type SqliteRepository interface {
	GetAllTasks() kb.Tasks
	ChangeState(state, id string) error
	SetAsDone(id string) error
	AddTicket(title, desc string) error
	ChangeTicket(state, id string)
}
