package server

import (
	"net/http"
)

func (s *Server) handleUserInput(w http.ResponseWriter, r *http.Request) {
	board := s.engine.GetBoard()

	switch r.FormValue("action") {
	case "move-ticket":
		s.handleMoveTicket(w, r, board)
		break
	case "add-ticket":
		s.handleAddTicket(w, r, board)
		break
	case "delete-ticket":
		s.handleDeleteTicket(w, r, board)
		break
	case "add-state":
		s.handleAddState(w, r)
		break
	}

	// Reset engine data
	s.engine.SetBoard(s.repo.GetBoard())

	s.engine.Render(w, r)
}
