package engine

import "kanbanBoard/server/templates"

// setupDummies returns a dummy set of tasks
// that can be feed into the page
func setupDummies() templates.Tasks {

	toDoTickets := templates.Tickets{
		templates.Ticket{"Ticket1", "Description of the task 1", 1},
		templates.Ticket{"Ticket2", "Description of the task 2", 2},
		templates.Ticket{"Ticket3", "Description of the task 3", 3},
	}

	inProgressTickets := templates.Tickets{
		templates.Ticket{"Ticket1 in progress", "Task 1 - currently in progress", 4},
		templates.Ticket{"Ticket2 in progress", "Task 2 - currently in progress", 5},
		templates.Ticket{"Ticket3 in progress", "Task 3 - currently in progress", 6},
	}

	doneTickets := templates.Tickets{
		templates.Ticket{"Ticket1 done", "Task 1 - already done", 7},
		templates.Ticket{"Ticket2d done", "Task 2 - already done", 8},
		templates.Ticket{"Ticket3d done", "Task 3 - already done", 9},
	}

	return templates.Tasks{toDoTickets, inProgressTickets, doneTickets}
}
