package repository

import (
	"fmt"
	kb "kanbanBoard"
	"strings"
)

// GetAllTasks pulls all tasks from the database and converts them
// into Tasks struct
func (r Repo) GetAllTasks() kb.Tasks {

	// Get all todos
	allTodos, _ := r.db.Query("SELECT title, desc, state, deadline, priority, id FROM tasks ORDER BY priority DESC ;")
	tasks := kb.Tasks{}
	var title, desc, state, deadline, id string
	var priority int

	for allTodos.Next() {
		err := allTodos.Scan(&title, &desc, &state, &deadline, &priority, &id)
		if err != nil {
			fmt.Println(err)
		}

		// Format deadline for displaying
		// Remove time
		deadlineNoTime := strings.Split(deadline, " ")[0]

		// Prepare to reverse
		deadlineSnipplets := strings.Split(deadlineNoTime, "-")
		newDeadline := ""
		for i := 2; i >= 0; i-- {
			/// Reverse
			newDeadline += deadlineSnipplets[i]
			if i > 0 {
				newDeadline += "-"
			}
		}

		// Distinguish between states
		switch state {
		case "todo":
			tasks.ToDo = append(tasks.ToDo, kb.Ticket{
				Title:       title,
				Description: desc,
				Deadline:    newDeadline,
				Priority:    priority,
				ID:          id,
			})
			break

		case "inprogress":
			tasks.InProgress = append(tasks.InProgress, kb.Ticket{
				Title:       title,
				Description: desc,
				Deadline:    newDeadline,
				Priority:    priority,

				ID: id,
			})
			break

		case "done":
			tasks.Done = append(tasks.Done, kb.Ticket{
				Title:       title,
				Description: desc,
				Deadline:    newDeadline,
				Priority:    priority,
				ID:          id,
			})
			break
		}
	}

	return tasks
}
