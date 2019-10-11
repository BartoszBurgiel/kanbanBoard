package repository

import "kanbanBoard/core"

// SqliteRepository handles a sqlite database
type SqliteRepository interface {
	GetBoard() core.Board
	PushTicketToTheDatabase(core.TicketElement) error
	RemoveTicket(ticketID string) error
}
