package main

import (
	"gorm.io/gorm"
	"time"
)

type VehicleType string

const (
	Bike    VehicleType = "bike"
	Scooter VehicleType = "scooter"
)

type Vehicle struct {
	gorm.Model
	Type         VehicleType    `json:"type"`
	Availability bool           `json:"availability"`
	Reservations []*Reservation `gorm:"foreignKey:VehicleID"`
}

type ReserveRequest struct {
	IdVehicle string
	IdUser    string
}

type ReserveResponse struct {
	OrderNumber string    `json:"order_number"`
	StartTime   time.Time `json:"start_time"`
}

type User struct {
	gorm.Model
	Name         string
	Reservations []*Reservation `gorm:"foreignKey:UserID"`
}

type Reservation struct {
	gorm.Model
	OrderNumber string
	StartTime   time.Time
	EndTime     time.Time
	VehicleID   uint
	UserID      uint
	User        User    `gorm:"foreignKey:UserID"`
	Vehicle     Vehicle `gorm:"foreignKey:VehicleID"`
}

type LocationUpdate struct {
	OrderNumber string    `json:"order_number"`
	StartTime   time.Time `json:"start_time"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Timestamp   int64     `json:"timestamp"`
}

type Config struct {
	BootstrapServers         string
	GroupId                  string
	ReservationRequestTopic  string
	ReservationResponseTopic string
	UpdateLocationTopic      string
}
