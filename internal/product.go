package internal

import (
	"log"

	database "github.com/muyisz/PracticalTraining/database/operation"
	"github.com/muyisz/PracticalTraining/handler/body"
)

func GetProductDetail(pid int) (body.GetProductRes, error) {
	product, err := database.GetProductByID(pid)
	if err != nil {
		log.Printf("[GetProductDetail] GetProductByID err,pid:%+v,err:%+v", pid, err)
		return body.GetProductRes{}, err
	}

	return body.GetProductRes{
		Name:    product.Name,
		Type:    product.Type,
		Color:   product.Color,
		Price:   product.Price,
		State:   product.State,
		Number:  product.Number,
		Img:     product.Img,
		ImgPath: product.ImgPath,
		Cid:     product.Cid,
	}, nil

}
