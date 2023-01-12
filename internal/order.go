package internal

import (
	"log"

	database "github.com/muyisz/PracticalTraining/database/operation"
	"github.com/muyisz/PracticalTraining/handler/body"
)

func GetUserOrder(user_name string) ([]body.OrderRes, error) {
	orderList := []body.OrderRes{}
	user, err := database.GetUserByUserName(user_name)
	if err != nil {
		log.Printf("[GetUserOrder] GetUserByUserName err,user_name:%+v,err:%+v", user_name, err)
		return nil, err
	}

	order, err := database.GetUserOrder(user.ID)
	if err != nil {
		log.Printf("[GetUserOrder] GetUserOrder err,oid:%+v,err:%+v", user.ID, err)
		return nil, err
	}
	for _, item := range order {
		orderList = append(orderList, body.OrderRes{
			ID:        item.ID,
			StartTime: item.Time,
			Status:    item.State,
		})
	}

	return orderList, nil
}

func CreateOrder(pid int, num int, address string, userName string) error {
	user, err := database.GetUserByUserName(userName)
	if err != nil {
		log.Printf("[CreateOrder] GetUserByUserName err,user_name:%+v,err:%+v", userName, err)
		return err
	}
	//todo 事务
	orderId, err := database.CreateOrder(user.ID, pid, address)
	if err != nil {
		log.Printf("[CreateOrder] CreateOrder err,userID:%+v,pid:%+v,address:%+v,err:%+v", user.ID, pid, address, err)
		return err
	}

	if err := database.CreateOrderDetail(orderId, pid, num); err != nil {
		log.Printf("[CreateOrder] CreateOrderDetail err,orderId:%+v,pid:%+v,num:%+v,err:%+v", orderId, pid, num, err)
		return err
	}
	return nil
}
