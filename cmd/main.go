package main

import (
	"kanbanBoard/server"
	"kanbanBoard/server/database"
	"net/http"
)

func main() {

	// Open Db connection
	repo := database.Open()

	s, err := server.NewServer(repo)
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8080", s)
}
