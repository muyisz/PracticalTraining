package database

import (
	"log"

	"github.com/muyisz/PracticalTraining/database/model"
	"github.com/muyisz/PracticalTraining/util"
)

func GetPasswordByUserName(userName string) (string, error) {
	user := model.Member{}
	if err := db.Where("name=?", userName).First(&user).Error; util.DBQueryErr(err) {
		log.Printf("[GetPasswordByUserName] First err,userName:%+v,err:%+v", userName, err)
		return "", err
	}
	return user.Password, nil
}