package repository

import (
	"fmt"

	"github.com/google/uuid"
)

// UpdateTicketState changes the state of a given ticket
func (r Repo) UpdateTicketState(stateID, id string) error {

	query, err := r.db.Prepare(`UPDATE tickets
								SET stateID = ? 
								WHERE id = ? ;`)
	if err != nil {
		fmt.Println(err)
	}

	res, err := query.Exec(stateID, id)
	n, _ := res.RowsAffected()
	fmt.Println("Updated", n, "rows")

	return err
}

// SetTicketAsDoneAndDelete sets given ticket as done and removes it
// from the database
func (r Repo) SetTicketAsDoneAndDelete(id string) error {

	query, _ := r.db.Prepare(`DELETE FROM tickets
								WHERE id = ? ;`)

	res, err := query.Exec(id)
	n, _ := res.RowsAffected()
	fmt.Println("Updated", n, "rows")

	return err
}

// AddNewTicket puts a ticket with given title and desc
// into the database
func (r Repo) AddNewTicket(title, desc, deadline string, priority int, state string) error {

	// Transfer data to inprogress
	query, err := r.db.Prepare(`INSERT INTO tickets
										VALUES (
												?, 
												?, 
												?, 
												?,
												?,
												?
											) ;`)

	fmt.Println(err)

	id := uuid.New().String()
	res, err := query.Exec(title, desc, deadline, priority, id, state)
	n, _ := res.RowsAffected()
	fmt.Println("Updated", n, "rows")

	return err
}
