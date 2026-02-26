package services

type NoteService struct{
	data []Note
}

type Note struct{
	Id int
	Name string
}

func (n *NoteService) GetNotes() []Note{
	return n.data
}

func (n *NoteService) CreateNote(id int, name string) int{
	temp := Note{
		Id: id,
		Name: name,
	}
	n.data = append(n.data, temp)
	return temp.Id
}