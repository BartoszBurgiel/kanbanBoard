package repository

import (
	"fmt"
	kb "kanbanBoard"
)

// GetBoard pulls all tasks from the database and converts them
// into Tasks struct
func (r Repo) GetBoard() kb.Board {

	board := kb.Board{}

	stateMap := make(map[string]*kb.State)

	var stateName, stateID, title, desc, deadline, id, ticketStateID string
	var priority, limit, position int

	// Get all states
	allStates, err := r.db.Query(`SELECT * FROM states ORDER BY states.position ASC; `)
	if err != nil {
		fmt.Println(err)
	}

	for allStates.Next() {

		err := allStates.Scan(&stateName, &stateID, &limit, &position)
		if err != nil {
			fmt.Println(err)
		}

		// Put id and state name to map
		stateMap[stateID] = &kb.State{
			State:    stateName,
			ID:       stateID,
			Position: position,
			Limit:    limit,
		}

	}

	// Get tickets
	tickets, err := r.db.Query("SELECT * FROM tickets ORDER BY tickets.priority DESC;")
	if err != nil {
		fmt.Println(err)
	}

	for tickets.Next() {

		err := tickets.Scan(&title, &desc, &deadline, &priority, &id, &ticketStateID)
		if err != nil {
			fmt.Println(err)
		}

		// Add ticket to current ticketstate
		stateMap[ticketStateID].Tickets = append(stateMap[ticketStateID].Tickets, kb.Ticket{
			Title:       title,
			Description: desc,
			Deadline:    deadline,
			Priority:    priority,
			ID:          id,
		})
	}

	// Assemble board
	for _, v := range stateMap {
		board.States = append(board.States, *v)
	}

	// Sort board
	for i := 0; i < len(board.States)-1; i++ {

		// Check priority
		if board.States[i].Position > board.States[i+1].Position {

			j := i
			for j < len(board.States)-1 {

				if board.States[j].Position > board.States[j+1].Position {
					// Swap
					temp := board.States[j+1]
					board.States[j+1] = board.States[j]
					board.States[j] = temp
					j++
				}
			}
			i = 0
		}
		fmt.Print(".")
	}

	fmt.Println(board)
	return board
}
