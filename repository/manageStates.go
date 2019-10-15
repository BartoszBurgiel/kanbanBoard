package repository

import (
	"fmt"
	"kanbanBoard/core"

	"github.com/google/uuid"
)

// AddNewState pushes a new state into the database
func (r Repo) AddNewState(name string, position, limit int) error {

	res, err := r.db.Exec("INSERT INTO states VALUES(?, ?, ?, ?) ; ", name, uuid.New().String(), limit, position)
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

// PushStateToTheDatabase updates a given state in the database
func (r Repo) PushStateToTheDatabase(s core.State) error {

	// Check if state is in the database
	row, err := r.db.Query("SELECT COUNT(states.stateID) FROM states WHERE states.stateID = ? ; ", s.ID)
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
		res, err := r.db.Exec("UPDATE states SET name = ?, ticket_limit = ?, position = ? WHERE stateID = ? ;", s.State, s.Limit, s.Position, s.ID)
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
		res, err := r.db.Exec("INSERT INTO states VALUES(?, ?, ?, ?) ;", s.State, s.ID, s.Limit, s.Position)
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

// RemoveState deletes a state with a given ID from the database
func (r Repo) RemoveState(stateID string) error {
	res, err := r.db.Exec("DELETE FROM states WHERE stateID = ? ", stateID)
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

	// Remove all belonging tickets as well
	resTickets, errTickets := r.db.Exec("DELETE FROM tickets WHERE stateID = ? ", stateID)
	if errTickets != nil {
		fmt.Println(errTickets)
		return errTickets
	}

	nTickets, errTickets := resTickets.RowsAffected()
	if errTickets != nil {
		fmt.Println(errTickets)
		return errTickets
	}

	fmt.Println("Updated", nTickets, "rows")

	return nil
}
