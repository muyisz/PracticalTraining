package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muyisz/PracticalTraining/constant"
	"github.com/muyisz/PracticalTraining/handler/body"
	"github.com/muyisz/PracticalTraining/internal"
	"github.com/muyisz/PracticalTraining/util"
)

func GetOrder(c *gin.Context) {
	userName, err := c.Cookie("user_name")
	if err != nil {
		log.Printf("[GetOrder] Cookie err,err:%+v", err)
		c.JSON(http.StatusOK, body.Res{
			Code: -1,
			Msg:  constant.GetCookieFailed,
		})
		return
	}
	orderList, err := internal.GetUserOrder(userName)
	if err != nil {
		log.Printf("[GetOrder] GetUserOrder err,userName:%+v,er:%+v", userName, err)
		log.Printf("[GetOrder] Cookie err,err:%+v", err)
		c.JSON(http.StatusOK, body.Res{
			Code: -1,
			Msg:  constant.GetOrderFailed,
		})
		return
	}
	c.JSON(http.StatusOK, body.Res{
		Code: 0,
		Data: orderList,
	})
}

func CreateOrder(c *gin.Context) {
	userName, err := c.Cookie("user_name")
	if err != nil {
		log.Printf("[CreateOrder] Cookie err,err:%+v", err)
		c.JSON(http.StatusOK, body.Res{
			Code: -1,
			Msg:  constant.GetCookieFailed,
		})
		return
	}

	parms, err := c.GetRawData()
	if err != nil {
		log.Printf("[CreateOrder] GetRawData err,reqBody:%+v,err:%+v", c.Request.Body, err)
		c.JSON(http.StatusBadRequest, body.Res{
			Code: -1,
		})
	}
	order := body.OrderReq{}
	if err := util.PostForm(parms, &order); err != nil {
		log.Printf("[CreateOrder] PostForm err,order:%+v,err:%+v", order, err)
		c.JSON(http.StatusBadRequest, body.Res{
			Code: -1,
		})
		return
	}

	if err := internal.CreateOrder(order.Pid, order.Number, order.Address, userName); err != nil {
		log.Printf("[CreateOrder] CreateOrder err,order:%+v,err:%+v", order, err)
		c.JSON(http.StatusBadRequest, body.Res{
			Code: -1,
		})
		return
	}
	c.JSON(http.StatusBadRequest, body.Res{
		Code: 0,
	})
}
