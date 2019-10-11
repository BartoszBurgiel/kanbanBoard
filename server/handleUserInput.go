package server

import (
	"fmt"
	"net/http"
	"strconv"
)

func (s *Server) handleUserInput(w http.ResponseWriter, r *http.Request) {

	if r.FormValue("ticketID") != "" {
		ticketID := r.FormValue("ticketID")
		ticktetState := r.FormValue("currentState")
		destinationID := r.FormValue("destination")

		fmt.Println("id", ticketID)
		fmt.Println("ticketState", ticktetState)
		fmt.Println("destinationID", destinationID)

		// Check if update is legal
		if ok, _ := s.repo.CheckStateLimit(destinationID); ok {
			s.repo.UpdateTicketState(destinationID, ticketID)
		} else {
			fmt.Println("LIMIT REACHED NO MORE TICKETS FOR THIS STATE! \nGET SOME STUFF DONE!")
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
