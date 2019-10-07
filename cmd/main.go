package main

import (
	"webserver/server/database"
)

func main() {

	//http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./ui/"))))

	database.InitRepository()

	// s, err := server.NewServer()
	// if err != nil {
	// 	panic(err)
	// }

	// http.ListenAndServe(":8080", s)
}
