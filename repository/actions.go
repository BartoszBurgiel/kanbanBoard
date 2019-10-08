package repository

import "fmt"

// ChangeTicket handles a given ticket according to it's state
func (r Repo) ChangeTicket(state, id string) {

	switch state {
	case "ToDo":
		err := r.ChangeState("inprogress", id)

		if err != nil {
			fmt.Println(err)
		}

		break
	case "InProgress":
		err := r.ChangeState("done", id)

		if err != nil {
			fmt.Println(err)
		}

		break
	case "Done":
		err := r.SetAsDone(id)

		if err != nil {
			fmt.Println(err)
		}

		break
	}

}
