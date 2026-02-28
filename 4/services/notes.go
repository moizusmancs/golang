package services

import (
	"fmt"

	internal "github.com/moizusmancs/students-api/internal/models"
	"gorm.io/gorm"
)

type NoteService struct{
	// data []Note
	db *gorm.DB
}

func (n *NoteService) InitService(db *gorm.DB){
	n.db = db
	n.db.AutoMigrate(&internal.NotesModel{})
}

// type Note struct{
// 	Id int
// 	Name string
// }

func (n *NoteService) GetNotes() []internal.NotesModel{
	// return n.data
	return []internal.NotesModel{}
}

func (n *NoteService) CreateNote(id int, title string, status string) {


	err := n.db.Create(&internal.NotesModel{
		Title: title,
		Status: status,
		},
	)
	if err != nil{
		fmt.Println(err)
	}
	// n.data = append(n.data, temp)
	// return noteToInsert.Id
}