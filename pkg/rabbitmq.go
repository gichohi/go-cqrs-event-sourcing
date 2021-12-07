package rabbitmq

import (
	"github.com/streadway/amqp"
	"os"
)

func Publish(msg string, queue string, exchange string) error {
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer func(channelRabbitMQ *amqp.Channel) {
		err := channelRabbitMQ.Close()
		if err != nil {

		}
	}(channelRabbitMQ)

	_, err = channelRabbitMQ.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msg),
	}

	if err := channelRabbitMQ.Publish(
		exchange,
		queue,
		false,
		false,
		message,
	); err != nil {
		return err
	}

	return err
}