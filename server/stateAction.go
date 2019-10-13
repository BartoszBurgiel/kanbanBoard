package server

import (
	"fmt"
	"net/http"
	"strconv"
)

// editState updates a given state
func (s *Server) handleEditState(w http.ResponseWriter, r *http.Request) {

	// Fetch variables
	newName := r.FormValue("newName")
	newLimit := r.FormValue("newLimit")
	newPosition := r.FormValue("newPosition")
	stateID := r.FormValue("stateID")

	fmt.Println("State edition:")
	fmt.Println("newName:", newName)
	fmt.Println("newLimit:", newLimit)
	fmt.Println("newPosition:", newPosition)
	fmt.Println("stateID:", stateID)

	// Convert limit and position to int
	nLimit, err := strconv.Atoi(newLimit)
	if err != nil {
		fmt.Println(err)
	}

	nPosition, err :=strconv.Atoi(newPosition)
	if err != nil {
		fmt.Println(err)
	}

	board := s.repo.GetBoard()

	for i := 0; i < len(board.States); i++ {

		// Search state with the same id
		if stateID == board.States[i].ID {
			// Replace
			board.States[i].State = newName
			board.States[i].Position = nPosition
			board.States[i].Limit = nLimit

			// Put changes to the databse
			s.repo.PushStateToTheDatabase(board.States[i])
			s.engine.SetBoard(board)
			break
		}
	}

}

// deleteState removes a given state from the database
func (s *Server) handleDeleteState() {

}

func (s *Server) handleAddState(w http.ResponseWriter, r *http.Request) {
	// Fetch variables
	newName := r.FormValue("newName")
	newLimit := r.FormValue("newLimit")
	newPosition := r.FormValue("newPosition")

	nPosition, err := strconv.Atoi(newPosition)
	if err != nil {
		fmt.Println(err)
	}

	nLimit, err := strconv.Atoi(newLimit)
	if err != nil {
		fmt.Println(err)
	}

	// Insert given info to the database
	error := s.repo.AddNewState(newName, nPosition, nLimit)
	if error != nil {
		fmt.Println(error)
	}
}
