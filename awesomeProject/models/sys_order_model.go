package models

import "gorm.io/gorm"

type SysOrder struct {
	gorm.Model
	CityName string `json:"city_name" gorm:"column:city_name"`
	PhoneNum int    `json:"phoneNum" gorm:"column:phone_num"`
}
