package engine

import "kanbanBoard/server/templates"

// setupDummies returns a dummy set of tasks
// that can be feed into the page
func setupDummies() templates.Tasks {

	toDoTickets := templates.Tickets{
		templates.Ticket{Title: "Ticket1", Description: "Description of the task 1", ID: 1},
		templates.Ticket{Title: "Ticket2", Description: "Description of the task 2", ID: 2},
		templates.Ticket{Title: "Ticket3", Description: "Description of the task 3", ID: 3},
	}

	inProgressTickets := templates.Tickets{
		templates.Ticket{Title: "Ticket1 in progress", Description: "Task 1 - currently in progress", ID: 4},
		templates.Ticket{Title: "Ticket2 in progress", Description: "Task 2 - currently in progress", ID: 5},
		templates.Ticket{Title: "Ticket3 in progress", Description: "Task 3 - currently in progress", ID: 6},
	}

	doneTickets := templates.Tickets{
		templates.Ticket{Title: "Ticket1 done", Description: "Task 1 - already done", ID: 7},
		templates.Ticket{Title: "Ticket2d done", Description: "Task 2 - already done", ID: 8},
		templates.Ticket{Title: "Ticket3d done", Description: "Task 3 - already done", ID: 9},
	}

	return templates.Tasks{ToDo: toDoTickets, InProgress: inProgressTickets, Done: doneTickets}
}
