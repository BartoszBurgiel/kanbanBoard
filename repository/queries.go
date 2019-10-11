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

// GetStateLimit returns the limit of the given state
// if the number of the tickets > GetStateLimit no more ticket can be
// moved to this state
func (r Repo) GetStateLimit(stateID string) (int, error) {

	// Get limit result set
	res, err := r.db.Query("SELECT states.limit FROM states WHERE states.stateID = ? LIMIT 1;", stateID)
	if err != nil {
		fmt.Println(err)
	}

	var n int

	for res.Next() {
		err := res.Scan(&n)
		if err != nil {
			return n, nil
		}
		return -1, err
	}
	return -1, err
}

// CheckStateLimit return true if a ticket can be pushed into the state
// -> if ticketCount+1 < stateLimit
func (r Repo) CheckStateLimit(stateID string) (bool, error) {

	limit, err := r.GetStateLimit(stateID)
	if err != nil {
		fmt.Println(err)
	}
	// Get the ticketCount
	var ticketCount int
	res, err := r.db.Query("SELECT COUNT(tickets.stateID) FROM tickets WHERE tickets.stateID = ? LIMIT 1 ;", stateID)
	if err != nil {
		fmt.Println(err)
	}

	for res.Next() {
		res.Scan(&ticketCount)

		return (ticketCount+1 > limit), nil
	}

	return false, err
}
