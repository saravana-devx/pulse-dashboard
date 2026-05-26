package rabbitmq

import (
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"pulseDashboard/internal/config"
)

// Connection retry settings. RabbitMQ takes ~20s to boot and its AMQP listener
// on 5672 comes up shortly after the node reports healthy, so the first few
// dials can be refused. Retry with a fixed delay instead of dying immediately.
const (
	connectMaxAttempts = 30
	connectRetryDelay  = 2 * time.Second
)

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func getRabbitMQConnURL() string {
	cfg := config.Get()
	return fmt.Sprintf("amqp://%s:%s@%s:5672/%s",
		cfg.RabbitMQUser,
		cfg.RabbitMQPass,
		cfg.RabbitMQHost,
		cfg.RabbitMQVHost,
	)
}

// function to create a new rabbitmq connection and channel
func NewRabbitMQConnection() *RabbitMQ {

	url := getRabbitMQConnURL()

	// connect to RabbitMQ, retrying while the broker finishes starting up
	var conn *amqp.Connection
	var err error
	for attempt := 1; attempt <= connectMaxAttempts; attempt++ {
		conn, err = amqp.Dial(url)
		if err == nil {
			break
		}
		log.Printf("RabbitMQ not ready (attempt %d/%d): %s", attempt, connectMaxAttempts, err)
		time.Sleep(connectRetryDelay)
	}
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ after %d attempts: %s", connectMaxAttempts, err)
	}

	// Open a RabbitMQ Channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a RabbitMQ channel : %s", err)
	}

	// Store RabbitMQ Connection and Channel
	return &RabbitMQ{
		Conn:    conn,
		Channel: ch,
	}
}

func (r *RabbitMQ) close() {
	if r.Channel != nil {
		if err := r.Channel.Close(); err != nil {
			log.Printf("Error closing RabbitMQ channel: %s", err)
		}
	}
	if r.Conn != nil {
		if err := r.Conn.Close(); err != nil {
			log.Printf("Error closing RabbitMQ connection: %s", err)
		}
	}
}
