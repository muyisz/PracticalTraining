package model

type Member struct {
	ID       int     `json:"id" gorm:"id"`
	Username string  `json:"username" gorm:"username"`
	Password string  `json:"password" gorm:"password"`
	Name     string  `json:"name" gorm:"name"`
	Identity string  `json:"identity" gorm:"identity"`
	Tel      string  `json:"tel" gorm:"tel"`
	Email    string  `json:"email" gorm:"email"`
	Balance  float64 `json:"balance" gorm:"balance"`
	State    int     `json:"state" gorm:"state"`
}
