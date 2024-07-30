package main

import (
	"fmt"
	_ "fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func ListVehicles(c *gin.Context) {
	var vehicles []Vehicle
	if err := db.Find(&vehicles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during load vehicles."})
		return
	}
	c.JSON(http.StatusOK, vehicles)
}

func GetVehicle(c *gin.Context) {
	id := c.Param("id")
	var vehicle Vehicle
	if err := db.First(&vehicle, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Could not find vehicle"})
		return
	}
	c.JSON(http.StatusOK, vehicle)
}

func AddVehicle(c *gin.Context) {
	var newVehicle Vehicle

	vehicleType := c.Query("vehicleType")
	newVehicle.Type = VehicleType(vehicleType)

	create(&newVehicle)
	c.JSON(http.StatusOK, newVehicle)
}

func ReserveVehicle(request ReserveRequest) {

	var reservation = createReservation(request)

	response := getReservationResponse(&reservation)

	fmt.Printf("sending response order number : " + response.OrderNumber)
	sendReserveResponse(response)
}

func createReservation(request ReserveRequest) Reservation {
	vehicle := Vehicle{}
	db.First(&vehicle, request.IdVehicle)
	user := User{}
	db.First(&user, request.IdUser)

	startTime := time.Now()
	orderNumber := fmt.Sprintf("order-%d", time.Now().UnixNano())

	vehicle.Availability = false
	db.Save(vehicle)

	reservation := Reservation{
		OrderNumber: orderNumber,
		StartTime:   startTime,
		EndTime:     startTime.Add(time.Hour),
		User:        user,
		Vehicle:     vehicle,
	}

	db.Create(&reservation)

	return reservation
}

func getReservationResponse(reservation *Reservation) ReserveResponse {
	startTime := time.Now()

	return ReserveResponse{
		OrderNumber: reservation.OrderNumber,
		StartTime:   startTime,
	}
}

func AddUser(c *gin.Context) {
	var newUser User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Adding user error: " + err.Error()})
		return
	}

	create(&newUser)
	c.JSON(http.StatusOK, newUser)
}

func create(value interface{}) error {
	result := db.Create(value)
	if result.Error != nil {
		log.Printf("Error saving to DB: %v\n", result.Error)
	}
	return result.Error
}
