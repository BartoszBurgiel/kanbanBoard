package repository

import (
	"fmt"
	kb "kanbanBoard"
)

// GetAllTasks pulls all tasks from the database and converts them
// into Tasks struct
func (r Repo) GetAllTasks() kb.Tasks {

	tasks := kb.Tasks{}

	// Get all columns and their tickets
	allColumns, err := r.db.Query(`SELECT columns.name, columns.columnId, 
											tickets.title, tickets.desc,  tickets.deadline, tickets.priority, tickets.id, tickets.columnID
									FROM columns 
									INNER JOINT tickets ON 
									tickets.columnID = columns.columnID ; `)
	if err != nil {
		fmt.Println(err)
	}

	var columnName, columnID, title, desc, deadline, id, columnIDTicket string
	var priority int

	for allColumns.Next() {

		err := allColumns.Scan(&columnName, &columnID, title, desc, deadline, priority, id, columnIDTicket)
		if err != nil {
			fmt.Println(err)
		}

		// Append column to the task
		tasks.Tickets = append(tasks.Tickets, kb.Tickets{
			Column: columnName,
			Tickets: []kb.Ticket{
				kb.Ticket{
					Title:       title,
					Description: desc,
					Deadline:    deadline,
					Priority:    priority,
					ID:          id,
				},
			},
		})

	}
	return tasks
}
