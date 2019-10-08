package server

import (
	"fmt"
	"kanbanBoard/engine"
	"net/http"
)

func (s *Server) handleUserInput(w http.ResponseWriter, r *http.Request) {

	if r.FormValue("ticketID") != "" {
		fmt.Println("Button pressed: ", r.FormValue("ticketID"))

		id := r.FormValue("ticketID")
		state := r.FormValue("state")

		s.repo.ChangeTicket(state, id)

	} else {
		fmt.Println("New ticket:")

		title := r.FormValue("newTitle")
		desc := r.FormValue("newDescription")

		fmt.Println("title: ", title)
		fmt.Println("desc: ", desc)

		s.repo.AddTicket(title, desc)

	}

	// Reset engine data
	s.engine.SetTasks(engine.RowsToTasks(s.repo.GetAllTasks()))

	s.engine.Render(w, r)

}
