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
		log.Printf("[GetOrder] GetUserOrder err,userName:%+v,err:%+v", userName, err)
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

func CancelOrder(c *gin.Context) {
	parms, err := c.GetRawData()
	if err != nil {
		log.Printf("[CancelOrder] GetRawData err,reqBody:%+v,err:%+v", c.Request.Body, err)
		c.JSON(http.StatusBadRequest, body.Res{
			Code: -1,
		})
	}
	reqBody := body.CancelOrderReq{}
	if err := util.PostForm(parms, &reqBody); err != nil {
		log.Printf("[CancelOrder] BindJSON err,reqBody:%+v,err:%+v", reqBody, err)
		c.JSON(http.StatusBadRequest, body.Res{
			Code: -1,
		})
		return
	}

	if err := internal.CancelOrder(reqBody.Oid); err != nil {
		log.Printf("[CancelOrder] CancelOrder err,orderID:%+v,err:%+v", reqBody.Oid, err)
		c.JSON(http.StatusBadRequest, body.Res{
			Code: -1,
		})
		return
	}
	c.JSON(http.StatusOK, body.Res{
		Code: 0,
	})

}
