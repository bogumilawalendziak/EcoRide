package main

import (
	"github.com/joho/godotenv"
	"os"
)

var config Config

func loadConfigFromEnv() Config {
	godotenv.Load()

	bootstrapServers := os.Getenv("BOOTSTRAP_SERVERS")
	groupId := os.Getenv("GROUP_ID")
	reserveReqTopic := os.Getenv("TOPIC_RESERVATION_REQUEST")
	reserveRespTopic := os.Getenv("TOPIC_RESERVATION_RESPONSE")
	updateLocationTopic := os.Getenv("TOPIC_LOCATION_UPDATE")
	return Config{
		BootstrapServers:         bootstrapServers,
		GroupId:                  groupId,
		ReservationRequestTopic:  reserveReqTopic,
		ReservationResponseTopic: reserveRespTopic,
		UpdateLocationTopic:      updateLocationTopic,
	}
}

func initKafka() {
	config = loadConfigFromEnv()
}
