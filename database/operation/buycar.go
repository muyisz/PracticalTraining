package database

import (
	"github.com/muyisz/PracticalTraining/constant"
	"github.com/muyisz/PracticalTraining/database/model"
	"github.com/muyisz/PracticalTraining/util"
)

func GetBuyCarByUserName(user_name string) ([]model.BuyCar, error) {
	buyCarList := []model.BuyCar{}

	if err := db.Table(constant.TableBuyCar).Where("username=?", user_name).Find(&buyCarList).Error; !util.DBQueryErr(err) {
		return nil, err
	}

	return buyCarList, nil
}
