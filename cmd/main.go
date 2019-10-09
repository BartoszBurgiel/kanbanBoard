package main

import (
	"fmt"
	"kanbanBoard/repository"
	"kanbanBoard/server"
	"net/http"
)

func main() {
	repo, err := repository.New("../repository/repo.db")
	if err != nil {
		fmt.Println(err)
	}

	s, err := server.NewServer(repo)
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8080", s)
}
