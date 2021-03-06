package queue

import (
	"fmt"
	"github.com/streadway/amqp"
	"os"
)

/*
set RABBITMQ_DEFAULT_USER=rabbitmq
set RABBITMQ_DEFAULT_PASS=rabbitmq
set RABBITMQ_DEFAULT_HOST=localhost
set RABBITMQ_DEFAULT_PORT=5672
set RABBITMQ_DEFAULT_VHOST=/
*/
func Connect() *amqp.Channel {

	dsn := "amqp://" + os.Getenv("RABBITMQ_DEFAULT_USER") + ":" + os.Getenv("RABBITMQ_DEFAULT_PASS") + "@" + os.Getenv("RABBITMQ_DEFAULT_HOST") + ":" + os.Getenv("RABBITMQ_DEFAULT_PORT") + os.Getenv("RABBITMQ_DEFAULT_VHOST")

	conn, err := amqp.Dial(dsn)
	if err != nil {
		panic(err.Error())
	}

	channel, err := conn.Channel()
	if err != nil {
		panic(err.Error())
	}
	return channel
}

func Notify(payload []byte, exchange string, routingKey string, ch *amqp.Channel) {
	err := ch.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(payload),
		},
	)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Message sent")
}
