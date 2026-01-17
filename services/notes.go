package services

import "strconv"

type Notesservice struct {
	notes []Note
}

type Note struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func NewNotesService() *Notesservice {
	return &Notesservice{
		notes: []Note{
			{Id: 1, Name: "Note 1"},
			{Id: 2, Name: "Note 2"},
		},
	}
}

func (n *Notesservice) GetNotes() []Note {
	return n.notes
}

func (n *Notesservice) CreateNotes(note Note) Note {
	n.notes = append(n.notes, note)
	return note
}

func (n *Notesservice) DeleteNotes(id string) Note {
	idInt, _ := strconv.Atoi(id)
	for i, note := range n.notes {
		if note.Id == idInt {
			deleted := note
			n.notes = append(n.notes[:i], n.notes[i+1:]...)
			return deleted
		}
	}
	return Note{}
}
