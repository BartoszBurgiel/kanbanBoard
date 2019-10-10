package server

import (
	"fmt"
	"net/http"
	"strconv"
)

func (s *Server) handleUserInput(w http.ResponseWriter, r *http.Request) {

	if r.FormValue("ticketID") != "" {
		fmt.Println("Button pressed: ", r.FormValue("ticketID"))

		id := r.FormValue("ticketID")

		// state logic
		if r.FormValue("state") != "" {
			s.repo.HandleTicketEvent(r.FormValue("state"), id)
		} else {

			// Check where ticket must be transfered
			switch r.FormValue("backState") {
			case "InProgress":
				s.repo.UpdateTicketState("todo", id)
				break
			case "Done":
				s.repo.UpdateTicketState("inprogress", id)
				break
			default:
				fmt.Println("State error!")
			}

		}

	} else {
		fmt.Println("New ticket:")

		title := r.FormValue("newTitle")
		desc := r.FormValue("newDescription")
		date := r.FormValue("newDate")
		priority := r.FormValue("newPriority")
		state := r.FormValue("newState")

		fmt.Println("title: ", title)
		fmt.Println("desc: ", desc)
		fmt.Println("date: ", date)
		fmt.Println("priority", priority)
		fmt.Println("state", state)

		// Convert priority to int
		newPriority, err := strconv.Atoi(priority)
		if err != nil {
			fmt.Println(err)
		}

		s.repo.AddNewTicket(title, desc, date, newPriority, state)

	}

	// Reset engine data
	s.engine.SetBoard(s.repo.GetBoard())

	s.engine.Render(w, r)
}
