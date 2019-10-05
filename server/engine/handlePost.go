package engine

import (
	"fmt"
	"net/http"
)

// handlePost handles a http request
func handlePost(r *http.Request) {
	switch r.Method {
	case "POST":

		r.ParseForm()

		if  r.FormValue("ticketID") != "" {
			fmt.Println("Button pressed: ", r.FormValue("ticketID"))
		} else {
			fmt.Println("New ticket:")
			fmt.Println("title: ", r.FormValue("newTitle"))
			fmt.Println("desc: ", r.FormValue("newDescription"))
		}

	default:

	}
}

func moveTicket(id string) {

}
