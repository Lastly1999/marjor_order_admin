package config

import (
	"awesomeProject/global"
	"awesomeProject/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func Setup() {
	//	构造连接池链接
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"marjor_user",
		"YH4zywwKkDZwj2sp",
		"127.0.0.1",
		"marjor_orders",
	)
	var err error
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	global.GLOBAL_DB, err = gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	global.GLOBAL_DB.Logger.LogMode(4)
	if err != nil {
		fmt.Errorf("the database connection failed may be due to a timeout")
		return
	}
	sqlDB, err := global.GLOBAL_DB.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	err = global.GLOBAL_DB.AutoMigrate(
		&models.SysOrder{},
	)
}
