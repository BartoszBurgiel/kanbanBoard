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

// PushStateToTheDatabase pushes a given state to the database
func (r Repo) PushStateToTheDatabase(s core.State) error {
	return nil
}

// RemoveState deletes a state with a given ID from the database
func (r Repo) RemoveState(stateID string) error {
	return nil
}