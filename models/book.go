package models

type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
}
