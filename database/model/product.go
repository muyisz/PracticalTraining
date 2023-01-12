package model

type Product struct {
	ID      int     `json:"id" gorm:"id"`
	Name    string  `json:"name" gorm:"name"`
	Type    string  `json:"type" gorm:"type"`
	Color   string  `json:"color" gorm:"color"`
	Price   float64 `json:"price" gorm:"price"`
	State   int     `json:"state" gorm:"state"`
	Number  int     `json:"number" gorm:"number"`
	Img     string  `json:"img" gorm:"img"`
	ImgPath string  `json:"img_path" gorm:"imgPath"`
	Cid     int     `json:"cid" gorm:"cid"`
}
