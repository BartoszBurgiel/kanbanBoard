package engine

import (
	"fmt"
	"net/http"
)

func handlePost(r *http.Request) {
	switch r.Method {
	case "POST":

		r.ParseForm()
		fmt.Println("Button pressed: ", r.FormValue("ticketID"))
	default:

	}
}

func moveTicket(id string) {

}
