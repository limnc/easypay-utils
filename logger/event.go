package logger

import (
	"encoding/json"
	"log"
	"time"

	"github.com/limnc/easypay-utils/rabbitmq"
)

type LoggerEvent struct {
	serviceName string
	rabbitMQ    *rabbitmq.RabbitMQ
}

// Register services initializes a logger event instance for a service
func RegisterService(serviceName string, rabbitMQ *rabbitmq.RabbitMQ) *LoggerEvent {
	return &LoggerEvent{
		serviceName: serviceName,
		rabbitMQ:    rabbitMQ,
	}
}

// Logaction logs an event and sends to the "logging" exchange
func (e *LoggerEvent) LogAction(logLevel, message string, metadata map[string]interface{}) {
	log_ := LogRequest{
		Service:    e.serviceName,
		LogLevel:   logLevel,
		LogMessage: message,
		CreatedAt:  time.Now().Format(time.RFC3339),
	}

	//If metadata exists, encode it to JSON
	if metadata != nil {
		metaJSON, err := json.Marshal(metadata)
		if err != nil {
			log.Printf("error marshalling metadata: %v\n", err)

		} else {
			log_.RequestBody = string(metaJSON)
		}
	}

	//Publish log to rabbitMQ
	err := e.rabbitMQ.Publish(rabbitmq.PublishMessage{
		Exchange:   "",
		RoutingKey: "logging",
		Body:       log_,
	})
	if err != nil {
		log.Printf("error publishing log message: %v\n", err)
	}
	//log.Println a message of time: service: service , Loglevel:loglevel, message and it will print to console
	log.Printf("time: %s, service: %s, logLevel: %s, message: %s\n", log_.CreatedAt, log_.Service, log_.LogLevel, log_.LogMessage)
}
