package logger

import (
	"encoding/json"
	"log"
	"time"

	"github.com/limnc/easypay-utils/rabbitmq"
)

// Producer interface to allow mocking in tests
type Producer interface {
	Publish(msg rabbitmq.PublishMessage) error
}

type LoggerEvent struct {
	serviceName string
	producer    Producer
}

// RegisterService initializes a LoggerEvent instance for a service
func RegisterService(serviceName string, producer Producer) *LoggerEvent {
	if producer == nil {
		log.Fatal("[Error] RegisterService: RabbitMQ instance is null")
	}

	logger := &LoggerEvent{
		serviceName: serviceName,
		producer:    producer,
	}

	log.Printf("[DEBUG] RegisterService: Service %s registered\n", serviceName)
	return logger
}

// LogAction logs an event and sends it to the "logging" queue
func (e *LoggerEvent) publishLog(level, message, metadata string) {
	// Publish log to RabbitMQ
	logEntry := LogRequest{
		Service:     e.serviceName,
		LogLevel:    level,
		LogMessage:  message,
		CreatedAt:   time.Now().Format(time.RFC3339),
		RequestBody: metadata,
	}
	err := e.producer.Publish(rabbitmq.PublishMessage{
		Exchange:   "",
		RoutingKey: "logging",
		Body:       logEntry,
	})

	if err != nil {
		log.Printf("error publishing log message: %v\n", err)
	}

	// Print log message to console
	log.Printf("time: %s, service: %s, logLevel: %s, message: %s\n",
		logEntry.CreatedAt, logEntry.Service, logEntry.LogLevel, logEntry.LogMessage)
}

func (e *LoggerEvent) LogInfo(message string, metadata any) {
	e.publishLog("INFO", message, handleMetaData(metadata))
}
func (e *LoggerEvent) LogWarning(message string, metadata any) {
	e.publishLog("WARNING", message, handleMetaData(metadata))
}
func (e *LoggerEvent) LogError(message string, metadata any) {
	e.publishLog("ERROR", message, handleMetaData(metadata))
}

func handleMetaData(metadata any) string {
	if metadata != nil {
		metaJSON, err := json.Marshal(metadata)
		if err != nil {
			log.Printf("Error handling the metadata: %v", err)
		} else {
			return string(metaJSON)
		}
	}
	return ""
}
