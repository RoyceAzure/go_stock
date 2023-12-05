package rabbitmq

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMqEmmiter struct {
	conn *amqp.Connection
}

func NewRabbitMqEmmiter(conn *amqp.Connection) *RabbitMqEmmiter {
	return &RabbitMqEmmiter{
		conn: conn,
	}
}

func (e *RabbitMqEmmiter) EmmitEvent(body string, contentType string) error {
	ch, err := e.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"testQ", // 队列名称
		false,   // 持久化
		false,   // 删除时是否自动
		false,   // 是否独占
		false,   // 是否阻塞
		nil,     // 其他属性
	)
	if err != nil {
		return err
	}

	err = ch.ExchangeDeclare(
		"testExchange", // 交换机名称
		"topic",        // 交换机类型：direct, fanout, topic, headers
		true,           // 是否持久化
		false,          // 是否自动删除
		false,          // 是否内部
		false,          // 是否等待服务器的响应
		nil,            // 其他参数
	)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %s", err)
	}

	err = ch.QueueBind(
		"testQ",        // 队列名称
		"#.test.#",     // 路由键
		"testExchange", // 交换机名称
		false,          // 是否等待服务器的响应
		nil,            // 其他参数
	)
	if err != nil {
		return err
	}

	err = ch.PublishWithContext(context.Background(),
		"testExchange", // 交换机
		"test",         // 路由
		false,          // 强制性
		false,          // 立即
		amqp.Publishing{
			ContentType: contentType,
			Body:        []byte(body),
		})
	if err != nil {
		return err
	}

	return nil
}
