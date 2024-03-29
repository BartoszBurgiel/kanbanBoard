package server

import (
	"fmt"
	"kanbanBoard/core"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

func (s *Server) handleMoveTicket(w http.ResponseWriter, r *http.Request, board *core.Board) {
	ticketID := r.FormValue("ticketID")
	ticktetState := r.FormValue("currentState")
	destinationID := r.FormValue("destination")

	fmt.Println("id", ticketID)
	fmt.Println("ticketState", ticktetState)
	fmt.Println("destinationID", destinationID)

	state, err := board.GetState(ticktetState)
	if err != nil {
		fmt.Println(err)
	}

	destinationState, err := board.GetState(destinationID)
	if err != nil {
		fmt.Println(err)
	}

	if destinationState.AllowsNewTicket() {

		// Update database
		changedTicket, err := state.GetTicket(ticketID)
		if err != nil {
			fmt.Println(err)
		}

		changedTicket.StateID = destinationID

		error := s.repo.PushTicketToTheDatabase(changedTicket)
		if error != nil {
			fmt.Println(error)
		}

		state.MoveTicketToNewState(destinationState, ticketID)

	} else {
		fmt.Println("LIMIT ERROR!")
	}
}

func (s *Server) handleAddTicket(w http.ResponseWriter, r *http.Request, board *core.Board) {
	fmt.Println("New ticket:")

	title := r.FormValue("newTitle")
	desc := r.FormValue("newDescription")
	date := r.FormValue("newDate")
	priority := r.FormValue("newPriority")
	stateID := r.FormValue("newState")

	fmt.Println("title: ", title)
	fmt.Println("desc: ", desc)
	fmt.Println("date: ", date)
	fmt.Println("priority", priority)
	fmt.Println("stateID", stateID)

	state, err := board.GetState(stateID)
	if err != nil {
		fmt.Println(err)
	}

	// Format prority
	priorityN, err := strconv.Atoi(priority)
	if err != nil {
		fmt.Println(err)
	}

	newTicket := core.TicketElement{
		Title:       title,
		Description: desc,
		Deadline:    date,
		Priority:    priorityN,
		ID:          uuid.New().String(),
		StateID:     stateID,
	}

	state.AddNewTicket(newTicket)

	error := s.repo.PushTicketToTheDatabase(newTicket)
	if error != nil {
		fmt.Println(error)
	}
}

func (s *Server) handleDeleteTicket(w http.ResponseWriter, r *http.Request, board *core.Board) {
	ticketID := r.FormValue("ticketID")
	stateID := r.FormValue("stateID")

	state, err := board.GetState(stateID)
	if err != nil {
		fmt.Println(err)
	}

	state.DeleteTicket(ticketID)

	// Remove from the database
	s.repo.RemoveTicket(ticketID)
}
