package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	initKafka()
	router := gin.Default()
	InitDB()
	InitializeAPI(router)

	go func() {
		fmt.Println("Starting ListenForReservationRequest")
		ListenForReservationRequest()
	}()
	go func() {
		fmt.Println("Starting ListenForLocationUpdate")
		ListenForLocationUpdate()
	}()

	router.Run(":8080")
}
