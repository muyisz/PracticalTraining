package model

type OrderDetail struct {
	ID     int `json:"id" gorm:"id"`
	Pid    int `json:"pid" gorm:"pid"`
	Oid    int `json:"oid" gorm:"oid"`
	Number int `json:"num" gorm:"num"`
}
