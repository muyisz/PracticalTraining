package database

import (
	"log"

	"github.com/muyisz/PracticalTraining/constant"
	"github.com/muyisz/PracticalTraining/database/model"
	"github.com/muyisz/PracticalTraining/util"
)

func GetOrderDetailByUserID(userID int) ([]model.OrderDetail, error) {
	orderDetailList := []model.OrderDetail{}

	if err := db.Table(constant.TableOrderDetail).Where("pid=?", userID).Find(&orderDetailList).Error; !util.DBQueryErr(err) {
		log.Printf("[GetOrderDetailByUserID] Find err,userID:%+v,err:%+v", userID, err)
		return nil, err
	}

	return orderDetailList, nil
}

func CreateOrderDetail(oid int, pid int, num int) error {
	orderDetail := model.OrderDetail{
		Pid:    pid,
		Oid:    oid,
		Number: num,
	}

	if err := db.Table(constant.TableOrderDetail).Create(orderDetail).Error; err != nil {
		log.Printf("[CreateOrderDetail] Create err,orderDetail:%+v,err:%+v", orderDetail, err)
		return err
	}

	return nil
}
