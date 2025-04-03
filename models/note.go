package models

// An example schematic for database tables
type Note struct {
	ID    int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Title string `json:"title" gorm:"size:255;not null"`
}
