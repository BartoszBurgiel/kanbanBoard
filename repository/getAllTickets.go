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
	states := []string{}

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

		states = append(states, stateName)

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

		tempStates := []string{}

		// Remove current state from states
		for _, s := range states {
			if s != stateMap[ticketStateID].State {
				tempStates = append(tempStates, s)
			}
		}

		// Add ticket to current ticketstate
		stateMap[ticketStateID].Tickets = append(stateMap[ticketStateID].Tickets, kb.TicketElement{
			Title:       title,
			Description: desc,
			Deadline:    deadline,
			Priority:    priority,
			ID:          id,
			StatesList:  tempStates,
		})
	}

	// Assemble board
	for _, v := range stateMap {
		board.States = append(board.States, *v)
	}

	fmt.Println(board)
	return board
}
