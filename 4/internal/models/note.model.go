package internal

type NotesModel struct {
	Id int `gorm:"primaryKey"`
	Title string `json:"title"`
	Status string `json:"status"`
}
