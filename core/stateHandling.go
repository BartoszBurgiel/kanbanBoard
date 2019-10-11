package core

import "fmt"

// AllowsNewTicket returns true if addidtion of a new ticket is legal
func (s State) AllowsNewTicket() bool {
	if s.Limit == 0 {
		return true
	}
	return (s.Limit > len(s.Tickets))
}

// MoveTicketToNewState transferres a ticket into given state
// and removes it from it's current state struct
func (s State) MoveTicketToNewState(newState *State, id string) {

	// Search for the needed ticket
	for n, t := range s.Tickets {
		if t.ID == id {

			// Change ticket's state attribute
			t.StateID = newState.ID

			// Append the ticket to the newState
			newState.Tickets = append(newState.Tickets, t)

			// Remove t from this state

			// Replace the given ticket with the last ticket in the list
			s.Tickets[n] = s.Tickets[len(s.Tickets)-1]

			// Trim ticket list
			s.Tickets = s.Tickets[:len(s.Tickets)-1]
			break
		}
	}
}

// AddNewTicket puts a ticket into the ticket list of a state
func (s State) AddNewTicket(t TicketElement) {
	s.Tickets = append(s.Tickets, t)
}

// DeleteTicket removes a ticket from the states
func (s State) DeleteTicket(id string) {
	// Search for the needed ticket
	for n, t := range s.Tickets {
		if t.ID == id {

			// Replace the given ticket with the last ticket in the list
			s.Tickets[n] = s.Tickets[len(s.Tickets)-1]

			// Trim ticket list
			s.Tickets = s.Tickets[:len(s.Tickets)-1]
			break
		}
	}
}

// GetTicket is a getter for the ticket in a state
func (s State) GetTicket(id string) (TicketElement, error) {
	for _, t := range s.Tickets {
		if t.ID == id {
			return t, nil
		}
	}

	return TicketElement{}, fmt.Errorf("No ticket found with id %s", id)
}
