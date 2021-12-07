package rabbitmq

import (
	"encoding/json"
	"github.com/gichohi/go-cqrs-rabbitmq/internal/dtos"
	store "github.com/gichohi/go-cqrs-rabbitmq/pkg/eventstore"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func consume(queue string, exchange string) {
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")

	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	defer func(connectRabbitMQ *amqp.Connection) {
		err := connectRabbitMQ.Close()
		if err != nil {

		}
	}(connectRabbitMQ)

	// Opening a channel to our RabbitMQ instance over
	// the connection we have already established.
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer func(channelRabbitMQ *amqp.Channel) {
		err := channelRabbitMQ.Close()
		if err != nil {

		}
	}(channelRabbitMQ)

	// Subscribing to QueueService1 for getting messages.
	messages, err := channelRabbitMQ.Consume(
		queue, // queue name
		"",              // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no local
		false,           // no wait
		nil,             // arguments
	)
	if err != nil {
		log.Println(err)
	}

	msgChannel := make(chan bool)

	go func() {
		for message := range messages {
			var productDto dtos.ProductDto

			err := json.Unmarshal(message.Body, &productDto)
			if err != nil {
				log.Fatal(err)
			}

			err = store.CreateEvent(&productDto)

			if err != nil {
				log.Printf(err.Error())
			}

			log.Printf(" > Received message: %s\n", message.Body)
		}
	}()

	<-msgChannel
}
