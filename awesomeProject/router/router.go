package router

import (
	"awesomeProject/api"
	"awesomeProject/middlewares/cors"
	"github.com/gin-gonic/gin"
)

func InitSysRoutes() *gin.Engine {
	app := gin.Default()
	app.Use(cors.Cors())
	makeOrderApi := api.MakeOrderApi{}
	app.POST("/register", makeOrderApi.RegisterMakeOrder)
	app.GET("/orders", makeOrderApi.FindMakeOrders)
	return app
}
