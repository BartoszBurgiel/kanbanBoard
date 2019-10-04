package engine

import (
	"fmt"
	"net/http"
)

func handlePost(r *http.Request) {
	switch r.Method {
	case "POST":
		//change data
		r.ParseForm()
		fmt.Println("Button pressed: ", r.Form["ticketID"])
	default:

	}
}
