package main

import (
	"net/http"
	"text/template"
	"webserver/server/templates"
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
						<div class="ticket-button">&rarr;</div>                       

					</div>
			
				{{end}}

				</div>
				<div class="board-body-column">

				{{range .InProgress}}
			
					<div class="ticket">
						<div class="ticket-header">{{.Title}}</div>
						<div class="ticket-desc">{{.Description}}</div>
						<div class="ticket-button">&rarr;</div>                       

					</div>
		
				{{end}}

				</div>
				<div class="board-body-column">

				{{range .Done}}
			
					<div class="ticket">
							<div class="ticket-header">{{.Title}}</div>
							<div class="ticket-desc">{{.Description}}</div>
							<div class="ticket-button">&rarr;</div>                       

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

	// mit post stucts manipulieren + action

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		toDoTickets := templates.Tickets{
			templates.Ticket{"T1", "T1desc"},
			templates.Ticket{"T2", "T2desc"},
			templates.Ticket{"T3", "T3desc"},
		}

		inProgressTickets := templates.Tickets{
			templates.Ticket{"T1ip", "T1descip"},
			templates.Ticket{"T2ip", "T2descip"},
			templates.Ticket{"T3ip", "T3descip"},
		}

		doneTickets := templates.Tickets{
			templates.Ticket{"T1d", "T1descd"},
			templates.Ticket{"T2d", "T2descd"},
			templates.Ticket{"T3d", "T3descd"},
		}

		tasks := templates.Tasks{toDoTickets, inProgressTickets, doneTickets}

		temp := template.Must(template.New("body").Parse(body))

		temp.Execute(w, tasks)
	})

	http.ListenAndServe(":8080", nil)
}
