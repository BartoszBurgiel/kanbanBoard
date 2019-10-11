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

	// stateProps stores all states and their IDs
	stateProps := []struct {
		Name string
		ID   string
	}{}

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

		// fill stateProps
		stateProps = append(stateProps, struct {
			Name string
			ID   string
		}{
			Name: stateName,
			ID:   stateID,
		})

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

		// Stores states only viable to the ticket
		tempStates := []struct {
			DestName string
			DestID   string
		}{}

		// Remove current state from states
		for _, s := range stateProps {

			// If not current ticket -> "remove"
			if s.Name != stateMap[ticketStateID].State {
				tempStates = append(tempStates, struct {
					DestName string
					DestID   string
				}{
					DestName: s.Name,
					DestID:   s.ID,
				})
			}
		}

		// Add ticket to current ticketstate
		stateMap[ticketStateID].Tickets = append(stateMap[ticketStateID].Tickets, kb.TicketElement{
			Title:       title,
			Description: desc,
			Deadline:    deadline,
			Priority:    priority,
			ID:          id,
			Destination: tempStates,
			StateID:     ticketStateID,
		})
	}

	// Assemble board
	for _, v := range stateMap {
		board.States = append(board.States, *v)
	}

	fmt.Println(board)
	return board
}

// PushTicketToTheDatabase pushes all structs into the database
func (r Repo) PushTicketToTheDatabase(t kb.TicketElement) error {

	// Check if a ticket is in the database
	row, err := r.db.Query("SELECT COUNT(tickets.id) FROM tickets WHERE tickets.id = ? ; ", t.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer row.Close()

	var n int

	// Get n
	for row.Next() {
		err := row.Scan(&n)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	query := " INTO tickets VALUES (?, ?, ?, ?, ?, ?) ;"
	prefix := ""

	if n != 0 {
		prefix = "REPLACE"
	} else {
		prefix = "INSERT"
	}

	res, err := r.db.Exec(prefix+query, t.Title, t.Description, t.Deadline, t.Priority, t.ID, t.StateID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	num, err := res.RowsAffected()
	if err != nil {
		fmt.Println("Updated", num, "rows")
		return err
	}

	return nil
}
