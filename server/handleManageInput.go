package server

import (
	"net/http"
)

func (s *Server) handleManageInput(w http.ResponseWriter, r *http.Request) {
	switch r.FormValue("action") {
	case "edit-state":
		s.handleEditState(w, r)

		break
	case "delete-state":
		s.handleDeleteState(w, r)

		break
	case "new-state":
		s.handleAddState(w, r)

		break
	}

	// Load page
	s.handleManageGet(w, r)
}
