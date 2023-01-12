package model

type Admin struct {
	UserName string `json:"user_name" gorm:"username"`
	Password string `json:"password" gorm:"password"`
}
