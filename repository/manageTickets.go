package repository

import (
	"fmt"
	"kanbanBoard/core"
)

// GetBoard pulls all tasks from the database and converts them
// into Tasks struct
func (r Repo) GetBoard() core.Board {

	board := core.Board{}
	stateMap := make(map[string]*core.State)

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
		stateMap[stateID] = &core.State{
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
		stateMap[ticketStateID].Tickets = append(stateMap[ticketStateID].Tickets, core.TicketElement{
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

	return board
}

// PushTicketToTheDatabase pushes all structs into the database
func (r Repo) PushTicketToTheDatabase(t core.TicketElement) error {

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

	fmt.Println("n : ", n)

	if n != 0 {
		fmt.Println("Here!!")
		fmt.Println(t.StateID, t.ID, "<= data")
		res, err := r.db.Exec("UPDATE tickets SET stateID = ? WHERE id = ? ;", t.StateID, t.ID)
		if err != nil {
			fmt.Println(err)
			return err
		}

		num, err := res.RowsAffected()
		fmt.Println("Updated", num, "rows")
		if err != nil {
			return err
		}
	} else {
		res, err := r.db.Exec("INSERT INTO tickets VALUES(?, ?, ?, ?, ?, ?) ;", t.Title, t.Description, t.Deadline, t.Priority, t.ID, t.StateID)
		if err != nil {
			fmt.Println(err)
			return err
		}

		num, err := res.RowsAffected()
		fmt.Println("Updated", num, "rows")
		if err != nil {
			return err
		}
	}

	return nil
}

// RemoveTicket deletes a ticket from the database
func (r Repo) RemoveTicket(t string) error {
	res, err := r.db.Exec("DELETE FROM tickets WHERE id = ? ", t)
	if err != nil {
		fmt.Println(err)
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Updated", n, "rows")
	return nil
}
