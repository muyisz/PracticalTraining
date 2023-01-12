package database

import (
	"log"

	"github.com/muyisz/PracticalTraining/constant"
	"github.com/muyisz/PracticalTraining/database/model"
	"github.com/muyisz/PracticalTraining/util"
)

func GetAdmin(userName string) (model.Admin, error) {
	admin := model.Admin{}

	if err := db.Table(constant.TableAdmin).Where("username=?", userName).Find(&admin).Error; !util.DBQueryErr(err) {
		log.Printf("[GetAdmin] Find err,userName:%+v,err:%+v", userName, err)
		return admin, err
	}

	return admin, nil
}
