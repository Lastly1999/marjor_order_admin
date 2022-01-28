package main

import (
	"awesomeProject/config"
	"awesomeProject/router"
)

func init() {
	config.Setup()
}

func main() {
	err := router.InitSysRoutes().Run(":3600")
	if err != nil {
		panic("server running error...")
		return
	}
}
