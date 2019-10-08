package main

import (
	"kanbanBoard/server"
	"kanbanBoard/server/database"
	"net/http"
)

func main() {)

	repo := database.NewRepo("../server/database/repository/repo.db")

	s, err := server.NewServer(repo)
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8080", s)
}
