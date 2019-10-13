package server

import (
	"fmt"
	"net/http"
)

func (s *Server) handleManageInput(w http.ResponseWriter, r *http.Request) {
	switch r.FormValue("action") {
	case "edit-state":

		// Fetch variables
		newName := r.FormValue("newName")
		newLimit := r.FormValue("newLimit")
		newPosition := r.FormValue("newPosition")
		stateID := r.FormValue("stateID")

		fmt.Println("State edition:")
		fmt.Println("newName:", newName)
		fmt.Println("newLimit:", newLimit)
		fmt.Println("newPosition:", newPosition)
		fmt.Println("stateID:", stateID)

		break
	case "delete-state":

		// Fetch variables
		stateID := r.FormValue("stateID")

		fmt.Println("State deletion:")
		fmt.Println("id: ", stateID)

		break
	case "new-state":


		//stateID := r.FormValue("stateID")

		fmt.Println("State addition:")
		fmt.Println("newName:", newName)
		fmt.Println("newLimit:", newLimit)
		fmt.Println("newPosition:", newPosition)

		s.handleAddState(w, r)

		break
	}

	// Load page
	s.handleManageGet(w, r)
}
