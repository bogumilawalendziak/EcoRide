package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := InitDB()
	InitializeAPI(router, db)
	router.Run(":8080")
}
