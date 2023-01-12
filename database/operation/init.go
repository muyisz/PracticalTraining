package database

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func InitDB() {
	once.Do(func() {
		dsn := "root:256275@tcp(localhost:3306)/practical_training?charset=utf8mb4&parseTime=True&loc=Local"
		var err error
		database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		db = database

		SetOrderIncrementID()
	})
}
