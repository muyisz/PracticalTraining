package database

import (
	"log"

	"github.com/muyisz/PracticalTraining/constant"
	"github.com/muyisz/PracticalTraining/database/model"
	"github.com/muyisz/PracticalTraining/util"
)

func GetProductByID(id int) (model.Product, error) {
	product := model.Product{}

	if err := db.Table(constant.TableProduct).Where("id=?", id).First(&product).Error; !util.DBQueryErr(err) {
		log.Printf("[GetProductByID] First err,id:%+v,err:%+v", id, err)
		return product, err
	}

	return product, nil
}
