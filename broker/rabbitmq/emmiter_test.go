package rabbitmq

import (
	"log"
	"testing"

	"github.com/RoyceAzure/go-stockinfo-broker/shared/util/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

func TestEmmitEvent(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	config := config.Config{
		RabbitMQAddress: "amqp://guest:guest@localhost:5672/",
	}
	conn, err := amqp.Dial(config.RabbitMQAddress)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	e := NewRabbitMqEmmiter(conn)
	e.EmmitEvent("{a:a}", "application/json")
}
