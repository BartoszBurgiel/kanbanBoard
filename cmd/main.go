package main

import (
	"kanbanBoard/server/engine"
	"net/http"
)

func main() {

	// Handle all CSS files
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("../server/html/style/"))))

	http.HandleFunc("/", engine.Run)

	http.ListenAndServe(":8080", nil)
}
