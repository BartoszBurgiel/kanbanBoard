package repository

import (
	"fmt"
	kb "kanbanBoard"
)

// GetBoard pulls all tasks from the database and converts them
// into Tasks struct
func (r Repo) GetBoard() kb.Board {

	board := kb.Board{}

	var stateName, stateID, title, desc, deadline, id, statesIDTicket string
	var priority, limit int

	// Get all states
	allStates, err := r.db.Query(`SELECT * FROM states ; `)
	if err != nil {
		fmt.Println(err)
	}

	for allStates.Next() {

		err := allStates.Scan(&stateName, &stateID, &limit)
		if err != nil {
			fmt.Println(err)
		}

		// Get tickets
		tickets, err := r.db.Query("SELECT * FROM tickets WHERE tickets.stateID = ? ; ", stateID)
		if err != nil {
			fmt.Println("s", err)
		}

		tempTickets := []kb.Ticket{}

		for tickets.Next() {

			err := tickets.Scan(&title, &desc, &deadline, &priority, &id, &statesIDTicket)
			if err != nil {
				fmt.Println(err)
			}

			// Append task to the temporary list
			tempTickets = append(tempTickets, kb.Ticket{
				Title:       title,
				Description: desc,
				Deadline:    deadline,
				Priority:    priority,
				ID:          id,
			})

		}

		// Append state to the board
		board.States = append(board.States, kb.State{
			State:   stateName,
			Tickets: tempTickets,
			Limit:   limit,
		})
	}

	return board
}
