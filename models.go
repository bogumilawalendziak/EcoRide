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
