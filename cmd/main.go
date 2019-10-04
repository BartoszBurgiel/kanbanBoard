package main

import (
	"fmt"
	"kanbanBoard/server/templates"
	"net/http"
	"text/template"
)

func main() {
	const body = `<!DOCTYPE html>
		<html>
			<head>
				<title>My kanban board</title>
				<link rel="stylesheet" href="/css/style.css" />
			</head>
		
		<body>
			<h1 id="header">My kanban board</h1>

			<div class="main">
				<div class="board-header">
					<div>Todo</div>
					<div>InProgress</div>
					<div>Done</div>
				</div>

			<div class="board-body">
				<div class="board-body-column">

				{{range .ToDo}}
			
					<div class="ticket">
						<div class="ticket-header">{{.Title}}</div>
						<div class="ticket-desc">{{.Description}}</div>
						<form action="/" method="POST">
							<input type="hidden" id="custId" name="ticketID" value="{{.ID}}">
							<input class="ticket-button" type="submit" name='{{.ID}}' value="&rarr;" />
						</form>                     

					</div>
			
				{{end}}

				</div>
				<div class="board-body-column">

				{{range .InProgress}}
			
					<div class="ticket">
						<div class="ticket-header">{{.Title}}</div>
						<div class="ticket-desc">{{.Description}}</div>
						<form action="/" method="POST">
							<input type="hidden" id="custId" name="ticketID" value="{{.ID}}">
							<input class="ticket-button" type="submit" name='{{.ID}}' value="&rarr;" />
						</form>                      

					</div>
		
				{{end}}

				</div>
				<div class="board-body-column">

				{{range .Done}}
			
					<div class="ticket">
							<div class="ticket-header">{{.Title}}</div>
							<div class="ticket-desc">{{.Description}}</div>
							<form action="/delete" method="post">
								<input type="hidden" id="custId" name="ticketID" value="{{.ID}}">
								<input class="ticket-button" type="submit" name='{{.ID}}' value="X" />
							</form>                       

					</div>

				{{end}}
				
				</div>

			<!-- close board-body -->
			</div>
	
		<!-- close main -->
		</div>
		</body>
		</html>`

	// Handle all CSS files
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("../server/html/style/"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

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

		switch r.Method {
		case "POST":
			//change data
			r.ParseForm()
			fmt.Println("Button pressed: ", r.Form["ticketID"])
		default:

		}

		tasks := templates.Tasks{toDoTickets, inProgressTickets, doneTickets}
		temp := template.Must(template.New("body").Parse(body))

		temp.Execute(w, tasks)

	})

	http.ListenAndServe(":8080", nil)
}
