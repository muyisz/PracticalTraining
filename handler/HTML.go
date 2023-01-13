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

func GetRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "Regist.html", gin.H{})
}

func GetAdminLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "backstage_login.html", gin.H{})
}

func GetCategoryPage(c *gin.Context) {
	c.HTML(http.StatusOK, "category_list.html", gin.H{})
}

func GetMessagePage(c *gin.Context) {
	c.HTML(http.StatusOK, "message_list.html", gin.H{})
}

func GetBackstageHome(c *gin.Context) {
	c.HTML(http.StatusOK, "backstage_home.html", gin.H{})
}

func GetOrderPage(c *gin.Context) {
	c.HTML(http.StatusOK, "order_list.html", gin.H{})
}

func GetProductListPage(c *gin.Context) {
	c.HTML(http.StatusOK, "product_list.html", gin.H{})
}

func GetUserListPage(c *gin.Context) {
	c.HTML(http.StatusOK, "user_list.html", gin.H{})
}

func GetMyOrderPage(c *gin.Context) {
	c.HTML(http.StatusOK, "my_order.html", gin.H{})
}
