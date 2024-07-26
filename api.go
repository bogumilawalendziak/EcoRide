package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializeAPI(router *gin.Engine, db *gorm.DB) {
	router.GET("/vehicles", ListVehicles)
	router.GET("/vehicle/:id", GetVehicle)
	router.POST("/vehicles", AddVehicle)
	router.POST("/reserve", ReserveVehicle)
	router.POST("/user", AddUser)
}
