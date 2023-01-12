package database

import (
	"log"

	"github.com/muyisz/PracticalTraining/constant"
	"github.com/muyisz/PracticalTraining/database/model"
	"github.com/muyisz/PracticalTraining/util"
)

func GetBuyCarByUserName(userName string) ([]model.BuyCar, error) {
	buyCarList := []model.BuyCar{}

	if err := db.Table(constant.TableBuyCar).Where("username=?", userName).Find(&buyCarList).Error; !util.DBQueryErr(err) {
		log.Printf("[GetBuyCarByUserName] Find err,userName:%+v,err:%+v", userName, err)
		return nil, err
	}

	return buyCarList, nil
}
