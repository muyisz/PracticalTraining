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

func CheckLoginStatus(c *gin.Context) {
	_, err := c.Cookie("user_name")
	if err != nil {
		c.HTML(http.StatusOK, "Login.html", gin.H{})
		c.Abort()
	}

	c.Next()
}

func Login(c *gin.Context) {
	parms, err := c.GetRawData()
	if err != nil {
		log.Printf("[Login] GetRawData err,reqBody:%+v,err:%+v", c.Request.Body, err)
		c.JSON(http.StatusBadRequest, body.Res{
			Code: -1,
		})
	}
	reqBody := body.LoginInfo{}
	if err := util.PostForm(parms, &reqBody); err != nil {
		log.Printf("[Login] BindJSON err,reqBody:%+v,err:%+v", reqBody, err)
		c.JSON(http.StatusBadRequest, body.Res{
			Code: -1,
		})
		return
	}
	login, err := internal.Login(reqBody.UserName, reqBody.Password)
	if err != nil || !login {
		log.Printf("[Login] internal.Login err,UserName:%+v,Password:%+v,err:%+v", reqBody.UserName, reqBody.Password, err)
		c.JSON(http.StatusOK, body.Res{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	c.SetCookie("user_name", reqBody.UserName, constant.CookieMaxAge, "/", constant.CookieAccessScope, false, false)
	c.JSON(http.StatusOK, body.Res{
		Code: 0,
	})
}
