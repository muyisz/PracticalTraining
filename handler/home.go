package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muyisz/PracticalTraining/constant"
	"github.com/muyisz/PracticalTraining/handler/body"
	"github.com/muyisz/PracticalTraining/internal"
)

func GetBuyCar(c *gin.Context) {
	userName, err := c.Cookie("user_name")
	if err != nil {
		log.Printf("[GetBuyCar] Cookie err,err:%+v", err)
		c.JSON(http.StatusOK, body.Res{
			Code: -1,
			Msg:  constant.GetCookieFailed,
		})
		return
	}
	buyCarRes, err := internal.GetUserBuyCar(userName)
	if err != nil {
		log.Printf("[GetBuyCar] GetUserBuyCar err,user_name:%+v,err:%+v", userName, err)
		c.JSON(http.StatusOK, body.Res{
			Code: -1,
			Msg:  constant.GetBuyCarFailed,
		})
		return
	}
	c.JSON(http.StatusOK, body.Res{
		Code: 0,
		Data: buyCarRes,
	})
}
