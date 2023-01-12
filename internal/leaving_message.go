package internal

import (
	"log"

	database "github.com/muyisz/PracticalTraining/database/operation"
)

func CreateLeavingMessage(title, msg, user_name string) error {
	user, err := database.GetUserByUserName(user_name)
	if err != nil {
		log.Printf("[CreateLeavingMessage] GetUserByUserName err,user_name:%+v,err:%+v", user_name, err)
		return err
	}

	err = database.CreateMessage(title, msg, user.ID)
	if err != nil {
		log.Printf("[CreateLeavingMessage] CreateMessage err,title:%+v,msg:%+v,err:%+v", title, msg, err)
		return err
	}

	return nil
}
