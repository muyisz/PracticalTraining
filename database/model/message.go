package model

import "time"

type Message struct {
	ID      int       `json:"id" gorm:"id"`
	Title   string    `json:"title" gorm:"title"`
	Content string    `json:"content" gorm:"content"`
	Date    time.Time `json:"date" gorm:"date"`
	Mid     int       `json:"mid" gorm:"mid"`
}
