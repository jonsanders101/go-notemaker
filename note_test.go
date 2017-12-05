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

func TestNotebook(t *testing.T) {
	notebook := notebook{}
	notebook.addNote("Favourite colours", "Magenta, Cerulean, Terrcotta")
	length := len(notebook.notes)
	note1 := notebook.notes[length-1]

	if note1.title != "Favourite colours" {
		t.Error("Expected a note to be in the notebook slice")
	}
}
