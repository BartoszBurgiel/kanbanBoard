package repository

import (
	kb "kanbanBoard"
)

// SqliteRepository handles a sqlite database
type SqliteRepository interface {
	GetBoard() kb.Board
	UpdateTicketState(newState, id string) error
	SetTicketAsDoneAndDelete(id string) error
	AddNewTicket(title, desc, deadline string, priority int, state string) error
	HandleTicketEvent(state, id string) error
}
