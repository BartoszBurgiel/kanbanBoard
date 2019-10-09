package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
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
		priority := r.FormValue("newPriority")

		fmt.Println("title: ", title)
		fmt.Println("desc: ", desc)
		fmt.Println("date: ", date)

		newDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			fmt.Println(err)
		}

		// Convert priority to int
		newPriority, err := strconv.Atoi(priority)
		if err != nil {
			fmt.Println(err)
		}

		s.repo.AddNewTicket(title, desc, newDate, newPriority)

	}

	// Clean database
	s.repo.ClearDatabase()

	// Reset engine data
	s.engine.SetTasks(s.repo.GetAllTasks())

	s.engine.Render(w, r)
}
