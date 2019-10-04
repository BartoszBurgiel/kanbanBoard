package main

import (
	"net/http"
	"text/template"
	"webserver/server/templates"
)

func main() {

	tmpl := template.Must(template.ParseFiles("../server/html/index.html"))

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

		tmpl.Execute(w, tasks)

	})

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("../server/html/style/"))))

	http.ListenAndServe(":8080", nil)
}
