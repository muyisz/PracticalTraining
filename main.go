package main

import (
	"github.com/gin-gonic/gin"
	database "github.com/muyisz/PracticalTraining/database/operation"
	"github.com/muyisz/PracticalTraining/router"
)

func main() {
	database.InitDB()

	r := gin.Default()
	router.InitRouter(r)

	r.Run(":8080")
}
