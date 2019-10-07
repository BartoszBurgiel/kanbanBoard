package server

import (
	"fmt"
	"net/http"
	"webserver/server/engine"
)

func (s *Server) handleUserInput(w http.ResponseWriter, r *http.Request) {

	//daten manipulieren
	//set/get foo bar bazz
	tasks := s.engine.GetTasks()

	if r.FormValue("ticketID") != "" {
		fmt.Println("Button pressed: ", r.FormValue("ticketID"))

		id := r.FormValue("ticketID")
		state := r.FormValue("state")

		var tickets *engine.Tickets
		var nextState *engine.Tickets

		switch state {
		case "ToDo":
			//tickets = &tasks.ToDo
			//nextState = &tasks.InProgress
			break
		case "InProgress":
			tickets = &tasks.InProgress
			nextState = &tasks.Done
			break
		case "Done":
			tickets = &tasks.Done
			nextState = nil
			break
		}

		// search in each slices for the given ID
		for i, ticket := range *tickets {
			if ticket.ID == id {

				fmt.Println("todo")
				fmt.Println("ticket: ", ticket)

				// modify

				*nextState = append(*nextState, ticket)

				if nextState == nil {
					(*tickets)[i] = (*tickets)[len(*tickets)-1]
					(*tickets) = (*tickets)[:len(*tickets)-1]

				}

				break
			}
		}

	} else {
		fmt.Println("New ticket:")

		title := r.FormValue("newTitle")
		desc := r.FormValue("newDescription")

		fmt.Println("title: ", title)
		fmt.Println("desc: ", desc)

		tasks.ToDo = append(tasks.ToDo, engine.NewTicket(title, desc))
	}

	s.engine.Render(w, r)

}
