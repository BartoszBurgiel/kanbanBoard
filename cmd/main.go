package main

import (
	"net/http"
	"webserver/server"
)

func main() {

	s, err := server.NewServer()
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8080", s)
}
