package kanbanboard

import "testing"

func TestState(t *testing.T) {
	state := State{}

	if state.ID != "" {
		t.Errorf("Expected state ID to be empty: %s", state.ID)
	}
	if !state.AllowsNewTicket() {
		t.Errorf("Exepcted AllowsNewTicket to be TRUE")
	}
}
