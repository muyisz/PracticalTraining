package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "Login.html", gin.H{})
}

func GetHome(c *gin.Context) {
	c.HTML(http.StatusOK, "Home.html", gin.H{})
}
