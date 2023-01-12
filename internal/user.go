package internal

import (
	"log"

	database "github.com/muyisz/PracticalTraining/database/operation"
	"github.com/muyisz/PracticalTraining/handler/body"
)

func GetUserBuyCar(user_name string) ([]body.BuyCarRes, error) {
	buyCarList, err := database.GetBuyCarByUserName(user_name)
	if err != nil {
		log.Printf("[GetUserBuyCar] GetBuyCarByUserName err,user_name:%+v,err:%+v", user_name, err)
		return nil, err
	}
	buyCarInfoList := []body.BuyCarRes{}

	for _, item := range buyCarList {
		product, err := database.GetProductByID(item.Pid)
		if err != nil {
			log.Printf("[GetUserBuyCar] GetProductByID err,id:%+v,err:%+v", item.Pid, err)
			continue
		}
		buyCarInfoList = append(buyCarInfoList, body.BuyCarRes{
			Image:    product.ImgPath,
			Describe: product.Name,
			Price:    product.Price,
			Number:   item.Num,
		})
	}

	return buyCarInfoList, nil
}

func CreateUser(userName string, password string, name string, identity string, tel string, email string, state int) error {

	if err := database.CreateUser(userName, password, name, identity, tel, email, state); err != nil {
		log.Printf("[CreateUser] CreateUser err,err:%+v", err)
		return err
	}

	return nil
}
