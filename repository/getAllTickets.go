package repository

import (
	"fmt"
	kb "kanbanBoard"
)

// GetBoard pulls all tasks from the database and converts them
// into Tasks struct
func (r Repo) GetBoard() kb.Board {

	board := kb.Board{}

	// Get all states and their tickets
	allStates, err := r.db.Query(`SELECT states.name, 
										tickets.title, tickets.desc, tickets.deadline, tickets.priority, tickets.id, tickets.statesID
										FROM states 
										INNER JOIN tickets
										ON states.statesID = tickets.statesID ; `)
	if err != nil {
		fmt.Println(err)
	}

	var stateName, title, desc, deadline, id, statesIDTicket string
	var priority int

	for allStates.Next() {

		err := allStates.Scan(&stateName, &title, &desc, &deadline, &priority, &id, &statesIDTicket)
		if err != nil {
			fmt.Println(err)
		}

		// If columns exist but 0 tickets
		// allColumns.

		// Append column to the task
		board.States = append(board.States, kb.State{
			State: stateName,
			Tickets: []kb.Ticket{
				kb.Ticket{
					Title:       title,
					Description: desc,
					Deadline:    deadline,
					Priority:    priority,
					ID:          id,
				},
			},
		})

	}

	fmt.Println(board)
	return board
}
