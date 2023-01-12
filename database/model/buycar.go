package model

type BuyCar struct {
	ID  int `json:"id" gorm:"id"`
	Pid int `json:"pid" gorm:"pid"`
	Num int `json:"num" gorm:"num"`
	Mid int `json:"mid" gorm:"mid"`
}
