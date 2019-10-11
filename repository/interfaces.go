package repository

import (
	kb "kanbanBoard"
)

// SqliteRepository handles a sqlite database
type SqliteRepository interface {
	GetBoard() kb.Board
	UpdateBoard(kb.Board) error
}
