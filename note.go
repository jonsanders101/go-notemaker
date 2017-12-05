package note

type note struct {
	title string
	body  string
}

type notebook struct {
	notes []note
}

func (notebook *notebook) addNote(newTitle string, newBody string) {
	newNote := note{title: newTitle, body: newBody}
	notebook.notes = append(notebook.notes, newNote)
}
