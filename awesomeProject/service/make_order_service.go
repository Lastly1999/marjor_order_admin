package service

import (
	"awesomeProject/global"
	"awesomeProject/models"
	"awesomeProject/models/request"
)

type MakeOrderService struct {
}

type IMakeOrderService interface {
	RegisterOrder(request *request.CreateOrderRequest) error
	FindMakeOrders() ([]models.SysOrder, error)
}

func (makeOrderService *MakeOrderService) RegisterOrder(request *request.CreateOrderRequest) error {
	var sysOrder models.SysOrder
	sysOrder.CityName = request.City
	sysOrder.PhoneNum = request.PhoneNum
	err := global.GLOBAL_DB.Create(&sysOrder).Error
	if err != nil {
		return err
	}
	return nil
}

func (makeOrderService *MakeOrderService) FindMakeOrders() ([]models.SysOrder, error) {
	var sysOrders []models.SysOrder
	err := global.GLOBAL_DB.Find(&sysOrders).Error
	if err != nil {
		return nil, err
	}
	return sysOrders, nil
}
