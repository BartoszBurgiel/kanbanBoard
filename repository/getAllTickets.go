package repository

import (
	"fmt"
	kb "kanbanBoard"
)

// GetBoard pulls all tasks from the database and converts them
// into Tasks struct
func (r Repo) GetBoard() kb.Board {

	board := kb.Board{}

	stateMap := make(map[string]struct {
		name  string
		limit int
	})

	var stateName, stateID, title, desc, deadline, id, statesIDTicket string
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
		stateMap[stateID] = struct {
			name  string
			limit int
		}{
			stateName,
			limit,
		}

	}

	// Get tickets
	tickets, err := r.db.Query("SELECT * FROM tickets ORDER BY tickets.stateID DESC;")
	if err != nil {
		fmt.Println(err)
	}

	tempState := kb.State{}
	newTempState := kb.State{}
	currentState := ""
	currentStateCheck := true
	newTempStateCheck := true

	for tickets.Next() {

		err := tickets.Scan(&title, &desc, &deadline, &priority, &id, &statesIDTicket)
		if err != nil {
			fmt.Println(err)
		}

		// Only in the first run define in the beginning of the loop
		if currentStateCheck {
			currentState = statesIDTicket
			currentStateCheck = false
		}

		fmt.Println("current state:", currentState)
		fmt.Println("ticekt state:", statesIDTicket)

		// check if currentState didn't change
		if currentState == statesIDTicket {

			// 'Update' only once
			if newTempStateCheck {
				tempState = newTempState
				newTempState = kb.State{}
				newTempStateCheck = false
			}

			tempState.State = stateMap[statesIDTicket].name
			tempState.Limit = stateMap[statesIDTicket].limit

			// Append tasks to the temporary list
			tempState.Tickets = append(tempState.Tickets, kb.Ticket{
				Title:       title,
				Description: desc,
				Deadline:    deadline,
				Priority:    priority,
				ID:          id,
			})

		} else {

			// New tempState for the current ticket (that's skipped)
			newTempState = kb.State{
				State: stateMap[statesIDTicket].name,
				Tickets: []kb.Ticket{kb.Ticket{
					Title:       title,
					Description: desc,
					Deadline:    deadline,
					Priority:    priority,
					ID:          id,
				}},
				Limit: stateMap[statesIDTicket].limit,
			}
			newTempStateCheck = true

			// Add to board
			board.States = append(board.States, tempState)

			// clear tempState
			tempState = kb.State{}
		}

		// redefine current state to used state
		currentState = statesIDTicket
	}

	// Add last items to board
	board.States = append(board.States, newTempState)

	fmt.Println(board)
	return board
}
