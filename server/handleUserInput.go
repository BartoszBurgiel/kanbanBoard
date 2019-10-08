package server

import (
	"fmt"
	"net/http"
)

func (s *Server) handleUserInput(w http.ResponseWriter, r *http.Request) {

	if r.FormValue("ticketID") != "" {
		fmt.Println("Button pressed: ", r.FormValue("ticketID"))

		id := r.FormValue("ticketID")
		state := r.FormValue("state")

		s.repo.HandleTicketEvent(state, id)

	} else {
		fmt.Println("New ticket:")

		title := r.FormValue("newTitle")
		desc := r.FormValue("newDescription")

		fmt.Println("title: ", title)
		fmt.Println("desc: ", desc)

		s.repo.AddNewTicket(title, desc)

	}

	// Reset engine data
	s.engine.SetTasks(s.repo.GetAllTasks())

	s.engine.Render(w, r)

}
