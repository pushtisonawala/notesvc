package services

type Notesservice struct{}
type Note struct{
	Id int `json:"id"`
	Name string `json:"name"`
}
func (n *Notesservice) GetNotes() []Note {
	data:=[]Note{
		{Id:1,Name:"Note 1"},
				{Id:2,Name:"Note 2"},

	}
	return data
}

func (n *Notesservice) CreateNotes() string {
	return "post request notes"
}
