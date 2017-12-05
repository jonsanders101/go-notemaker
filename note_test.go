package note

import "testing"

func TestNote(t *testing.T) {
	note := note{title: "Shopping", body: "3 x eggs"}

	noteTitle, noteBody := note.title, note.body

	if noteTitle != "Shopping" {
		t.Error("Did not instantiate with correct title")
	}

	if noteBody != "3 x eggs" {
		t.Error("Did not instantiate with correct body")
	}
}
