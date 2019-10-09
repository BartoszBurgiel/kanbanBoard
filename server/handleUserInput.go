package server

import (
	"fmt"
	"net/http"
	"strings"
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
		date := r.FormValue("newDate")

		fmt.Println("title: ", title)
		fmt.Println("desc: ", desc)

		// Reformat date
		dateSnipplets := strings.Split(date, "-")
		newDate := ""

		for i := 2; i >= 0; i-- {
			newDate += dateSnipplets[i]
			if i > 0 {
				newDate += "-"
			}
		}

		fmt.Println("date: ", newDate)
		s.repo.AddNewTicket(title, desc, newDate)

	}

	// Clean database
	s.repo.ClearDatabase()

	// Reset engine data
	s.engine.SetTasks(s.repo.GetAllTasks())

	s.engine.Render(w, r)

}
