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
