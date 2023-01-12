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

func PostLeavingMessage(c *gin.Context) {
	userName, err := c.Cookie("user_name")
	if err != nil {
		log.Printf("[PostLeavingMessage] Cookie err,err:%+v", err)
		c.JSON(http.StatusOK, body.Res{
			Code: -1,
			Msg:  constant.GetCookieFailed,
		})
		return
	}

	parms, err := c.GetRawData()
	if err != nil {
		log.Printf("[PostLeavingMessage] GetRawData err,reqBody:%+v,err:%+v", c.Request.Body, err)
		c.JSON(http.StatusBadRequest, body.Res{
			Code: -1,
			Msg:  constant.PostLeavingMessageFailed,
		})
	}
	reqBody := body.LeavingMessage{}
	if err := util.PostForm(parms, &reqBody); err != nil {
		log.Printf("[PostLeavingMessage] PostForm err,reqBody:%+v,err:%+v", reqBody, err)
		c.JSON(http.StatusBadRequest, body.Res{
			Code: -1,
			Msg:  constant.PostLeavingMessageFailed,
		})
		return
	}
	if err := internal.CreateLeavingMessage(reqBody.Title, reqBody.Msg, userName); err != nil {
		log.Printf("[PostLeavingMessage] CreateLeavingMessage err,title:%+v,msg:%+v,userName:%+v,err:%+v", reqBody.Title, reqBody.Msg, userName, err)
		c.JSON(http.StatusBadRequest, body.Res{
			Code: -1,
		})
		return
	}
	c.JSON(http.StatusOK, body.Res{
		Code: 0,
	})
}
