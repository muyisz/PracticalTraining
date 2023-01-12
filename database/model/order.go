package model

import "time"

type Order struct {
	ID      int       `json:"id" gorm:"id"`
	Time    time.Time `json:"time" gorm:"time"`
	State   int       `json:"state" gorm:"state"`
	Mid     int       `json:"mid" gorm:"mid"`
	Address string    `json:"address" gorm:"address"`
}
