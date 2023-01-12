package database

import (
	"log"
	"time"

	"github.com/muyisz/PracticalTraining/constant"
	"github.com/muyisz/PracticalTraining/database/model"
)

func CreateMessage(title string, Msg string, mid int) error {
	message := model.Message{
		Title:   title,
		Content: Msg,
		Date:    time.Now(),
		Mid:     mid,
	}
	if err := db.Table(constant.TableMessage).Create(&message).Error; err != nil {
		log.Printf("[CreateMessage] Create err,message:%+v,err:%+v", message, err)
		return err
	}

	return nil
}
