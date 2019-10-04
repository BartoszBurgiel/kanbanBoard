package engine

import (
	"html/template"
	"net/http"
)

// Run is the main function of the engine
// from here the webpage will be served and handled
func Run(w http.ResponseWriter, r *http.Request) {
	dummyData := setupDummies()

	// assemble the page
	temp := template.Must(template.New("body").Parse(body))

	handlePost(r)

	temp.Execute(w, dummyData)

}
