package main

import (
	"net/http"
	"webserver/server"
)

func main() {

	//http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./ui/"))))

	s, err := server.NewServer()
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8080", s)
}
