package main

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func sendReserveResponse(response ReserveResponse) error {
	producer := initKafkaProducer(config)
	fmt.Printf("Sending reservation response: %s\n", response)
	responseJson, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("Failed to marshal reservation response: %s\n", err)
		return err
	}

	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &config.ReservationResponseTopic, Partition: kafka.PartitionAny},
		Value:          responseJson,
	}, nil)

	if err != nil {
		fmt.Printf("Failed to produce message: %s\n", err)
		return err
	}
	fmt.Printf("Reservation response sent to \n" + config.ReservationResponseTopic)
	producer.Flush(15 * 1000)

	return nil
}

func ListenForReservationRequest() {
	consumer := initKafkaConsumer(config)
	fmt.Printf("Listening for reservation requests\n")

	var err = consumer.SubscribeTopics([]string{config.ReservationRequestTopic}, nil)
	if err != nil {
		fmt.Printf("Failed to subscribe to topics: %s", err)
		return
	}

	for {
		fmt.Printf("Listening for reservation requests message\n")
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			fmt.Printf("Consumer error: %s %s", err, msg)
			continue
		}

		var request ReserveRequest
		if err := json.Unmarshal(msg.Value, &request); err != nil {
			fmt.Printf("Failed to unmarshal ReserveRequest: %s", err)
			continue
		}

		ReserveVehicle(request)
	}
}

func ListenForLocationUpdate() {
	consumer := initKafkaConsumer(config)
	var err = consumer.SubscribeTopics([]string{config.UpdateLocationTopic}, nil)
	if err != nil {
		fmt.Printf("Failed to subscribe to topics: %s", err)
		return
	}

	for {
		fmt.Printf("reading message from location update topic\n")
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			fmt.Printf("Consumer error: %s %s", err, msg)
			continue
		}

		var response LocationUpdate
		if err := json.Unmarshal(msg.Value, &response); err != nil {
			fmt.Printf("Failed to unmarshal ReserveResponse: %s", err)
			continue
		}
		fmt.Printf("Starting location streaming for order %s\n, latitude: %s\n , longitude: %s\n", response.OrderNumber, response.Latitude, response.Longitude)

	}
}

func initKafkaProducer(config Config) *kafka.Producer {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": config.BootstrapServers})
	if err != nil {
		fmt.Printf("Failed to create Kafka producer: %s\n", err)
		return nil
	}
	return producer
}

func initKafkaConsumer(config Config) *kafka.Consumer {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.BootstrapServers,
		"group.id":          config.GroupId,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		fmt.Printf("Failed to create Kafka consumer: %s", err)
		return nil
	}

	fmt.Printf("**** Kafka consumer created\n")
	return consumer
}
