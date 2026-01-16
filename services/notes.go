package services

type Notesservice struct{}

func (n *Notesservice) GetNotes() string {
	return "Get request notes"
}

func (n *Notesservice) CreateNotes() string {
	return "post request notes"
}
