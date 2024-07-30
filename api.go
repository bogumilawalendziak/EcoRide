package main

import (
	"github.com/gin-gonic/gin"
)

func InitializeAPI(router *gin.Engine) {
	router.GET("/vehicles", ListVehicles)
	router.GET("/vehicle/:id", GetVehicle)
	router.POST("/vehicles", AddVehicle)
	router.POST("/user", AddUser)
}
