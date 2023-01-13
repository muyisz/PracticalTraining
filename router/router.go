package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muyisz/PracticalTraining/handler"
)

func InitRouter(router *gin.Engine) {
	router.LoadHTMLGlob("view/**/*.html")
	router.StaticFS("/css", http.Dir("./view/reception/css"))
	router.StaticFS("/images", http.Dir("./view/reception/images"))
	router.StaticFS("/js", http.Dir("./view/reception/js"))
	/////////////////////////////////////////////////////////
	router.StaticFS("/assets/css", http.Dir("./view/reception/assets/css"))
	router.StaticFS("/assets/js", http.Dir("./view/reception/assets/js"))
	////////////////////////////////////////////////////////
	router.GET("/", handler.GetHome)
	router.GET("/home", handler.GetHome)
	router.GET("/login", handler.GetLogin)
	router.POST("/login", handler.Login)
	router.GET("/buycar", handler.GetBuyCar)
	router.GET("/order", handler.GetOrder)
	router.POST("/leaving_message", handler.PostLeavingMessage)
	router.POST("/order", handler.CreateOrder)
	router.GET("/product_detail", handler.GetProductDetail)
	router.GET("/register", handler.GetRegister)
	router.POST("/register", handler.CreateUser)
	router.POST("/admin", handler.AdminLogin)
	router.GET("/admin", handler.GetAdminLogin)
	router.GET("/backstage_home", handler.GetBackstageHome)
	router.GET("/category_list", handler.GetCategoryPage)
	router.GET("/message_list", handler.GetMessagePage)
	router.GET("/order_list", handler.GetOrderPage)
	router.GET("/product_list", handler.GetProductListPage)
	router.GET("/user_list", handler.GetUserListPage)
	router.GET("/my_order", handler.GetMyOrderPage)
	router.POST("/cancel_order", handler.CancelOrder)
}
