package core

import "fmt"

// GetState returns a state from the board with the given index
func (b *Board) GetState(stateID string) (*State, error) {
	for _, s := range b.States {
		if s.ID == stateID {
			return &s, nil
		}
	}

	return &State{}, fmt.Errorf("No state found with id %s", stateID)
}
