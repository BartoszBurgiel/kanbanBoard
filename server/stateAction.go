package server

import (
	"fmt"
	"kanbanBoard/core"
	"net/http"
	"strconv"
)

// editState updates a given state
func (s *Server) handleEditState(st core.State) {

	// Replace the state from the board with the given state

	board := s.repo.GetBoard()

	for i := 0; i < len(board.States); i++ {

		// Search state with the same id
		if st.ID == board.States[i].ID {
			// Replace
			board.States[i] = st

			// Put changes to the databse
			s.repo.PushStateToTheDatabase(st)
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
