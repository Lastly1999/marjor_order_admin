package api

import (
	"awesomeProject/models/request"
	"awesomeProject/service"
	"github.com/gin-gonic/gin"
)

type MakeOrderApi struct {
}

type IMakeOrderApi interface {
	RegisterMakeOrder(ctx *gin.Context)
	FindMakeOrders(ctx *gin.Context)
}

var makeOrderService service.MakeOrderService

func (makeOrderApi MakeOrderApi) RegisterMakeOrder(ctx *gin.Context) {
	createOrderRequest := request.CreateOrderRequest{}
	err := ctx.ShouldBindJSON(&createOrderRequest)
	if err != nil {
		return
	}
	err = makeOrderService.RegisterOrder(&createOrderRequest)
	if err != nil {
		return
	}
	ctx.JSON(200, gin.H{
		"msg":  "success",
		"code": 200,
		"data": nil,
	})
}

func (makeOrderApi MakeOrderApi) FindMakeOrders(ctx *gin.Context) {
	orders, err := makeOrderService.FindMakeOrders()
	if err != nil {
		return
	}
	ctx.JSON(200, gin.H{
		"msg":  "success",
		"code": 200,
		"data": orders,
	})
}
