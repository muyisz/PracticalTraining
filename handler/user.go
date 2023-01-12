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

func CreateUser(c *gin.Context) {
	parms, err := c.GetRawData()
	if err != nil {
		log.Printf("[CreateUser] GetRawData err,reqBody:%+v,err:%+v", c.Request.Body, err)
		c.JSON(http.StatusBadRequest, body.Res{
			Code: -1,
		})
	}
	reqBody := body.UserInfo{}
	if err := util.PostForm(parms, &reqBody); err != nil {
		log.Printf("[CreateUser] BindJSON err,reqBody:%+v,err:%+v", reqBody, err)
		c.JSON(http.StatusBadRequest, body.Res{
			Code: -1,
		})
		return
	}

	if err := internal.CreateUser(reqBody.UserName, reqBody.Password, "", "", "", reqBody.Email, 0); err != nil {
		log.Printf("[CreateUser] CreateUser err,reqBody:%+v,err:%+v", reqBody, err)
		c.JSON(http.StatusOK, body.Res{
			Code: -1,
		})
		return
	}

	c.JSON(http.StatusOK, body.Res{
		Code: 0,
	})
}

func CheckLoginStatus(c *gin.Context) {
	_, err := c.Cookie("user_name")
	if err != nil {
		c.HTML(http.StatusOK, "Login.html", gin.H{})
		c.Abort()
	}

	c.Next()
}

func AdminLogin(c *gin.Context) {
	parms, err := c.GetRawData()
	if err != nil {
		log.Printf("[AdminLogin] GetRawData err,reqBody:%+v,err:%+v", c.Request.Body, err)
		c.JSON(http.StatusBadRequest, body.Res{
			Code: -1,
		})
	}
	reqBody := body.LoginInfo{}
	if err := util.PostForm(parms, &reqBody); err != nil {
		log.Printf("[AdminLogin] BindJSON err,reqBody:%+v,err:%+v", reqBody, err)
		c.JSON(http.StatusBadRequest, body.Res{
			Code: -1,
		})
		return
	}
	login, err := internal.AdminLogin(reqBody.UserName, reqBody.Password)
	if err != nil || !login {
		log.Printf("[AdminLogin] internal.AdminLogin err,UserName:%+v,Password:%+v,err:%+v", reqBody.UserName, reqBody.Password, err)
		c.JSON(http.StatusOK, body.Res{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, body.Res{
		Code: 0,
	})
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
