package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn *amqp.Connection
}

func NewRabbitMqConsumer(conn *amqp.Connection) *Consumer {
	return &Consumer{
		conn: conn,
	}
}

func (c *Consumer) ConsumeMsg() error {
	ch, err := c.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	msgs, err := ch.Consume(
		"testQ", // 队列
		"",      // 消费者
		true,    // 自动应答
		false,   // 独占
		false,   // 本地
		false,   // 阻塞
		nil,     // 其他属性
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
	return nil
}
