package logger

import (
	"encoding/json"
	"log"
	"time"

	"github.com/limnc/easypay-utils/rabbitmq"
)

// Define an interface to allow mocking in tests
type Publisher interface {
	Publish(msg rabbitmq.PublishMessage) error
}

type LoggerEvent struct {
	serviceName string
	publisher   Publisher
}

// RegisterService initializes a LoggerEvent instance for a service
func RegisterService(serviceName string, publisher Publisher) *LoggerEvent {
	if publisher == nil {
		log.Fatal("[Error] RegisterService: RabbitMQ instance is null")
	}

	logger := &LoggerEvent{
		serviceName: serviceName,
		publisher:   publisher,
	}

	log.Printf("[DEBUG] RegisterService: Service %s registered\n", serviceName)
	return logger
}

// LogAction logs an event and sends it to the "logging" queue
func (e *LoggerEvent) LogAction(logLevel, message string, metadata map[string]interface{}) {
	log_ := LogRequest{
		Service:    e.serviceName,
		LogLevel:   logLevel,
		LogMessage: message,
		CreatedAt:  time.Now().Format(time.RFC3339),
	}

	// If metadata exists, encode it to JSON
	if metadata != nil {
		metaJSON, err := json.Marshal(metadata)
		if err != nil {
			log.Printf("error marshalling metadata: %v\n", err)
		} else {
			log_.RequestBody = string(metaJSON)
		}
	}

	// Publish log to RabbitMQ
	err := e.publisher.Publish(rabbitmq.PublishMessage{
		Exchange:   "",
		RoutingKey: "logging",
		Body:       log_,
	})

	if err != nil {
		log.Printf("error publishing log message: %v\n", err)
	}

	// Print log message to console
	log.Printf("time: %s, service: %s, logLevel: %s, message: %s\n",
		log_.CreatedAt, log_.Service, log_.LogLevel, log_.LogMessage)
}
