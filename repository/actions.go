package repository

import "fmt"

// HandleTicketEvent handles a given ticket according to it's state
func (r Repo) HandleTicketEvent(state, id string) error {

	switch state {
	case "ToDo":
		err := r.UpdateTicketState("inprogress", id)

		if err != nil {
			fmt.Println(err)
			return err
		}

		break
	case "InProgress":
		err := r.UpdateTicketState("done", id)

		if err != nil {
			fmt.Println(err)
			return err
		}

		break
	case "Done":
		err := r.SetTicketAsDoneAndDelete(id)

		if err != nil {
			fmt.Println(err)
			return err
		}

		break
	}

	return nil
}
