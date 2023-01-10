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
	router.GET("/", handler.GetLogin)
	router.GET("/home", handler.GetHome)
	router.POST("/login", handler.Login)
}
