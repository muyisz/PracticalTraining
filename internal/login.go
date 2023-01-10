package internal

import (
	"fmt"
	"log"

	"github.com/muyisz/PracticalTraining/constant"
	database "github.com/muyisz/PracticalTraining/database/operation"
)

func Login(userName string, password string) (bool, error) {
	realPassword, err := database.GetPasswordByUserName(userName)
	if err != nil {
		log.Printf("[Login] GetPasswordByUserName err,userName:%+v,err:%+v", userName, err)
		return false, err
	}
	if realPassword != password {
		return false, fmt.Errorf(constant.WrongPassword)
	}
	return true, nil
}
