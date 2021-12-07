package commands

import (
	"encoding/json"
	"fmt"
	"github.com/gichohi/go-cqrs-rabbitmq/internal/models"
	"github.com/gichohi/go-cqrs-rabbitmq/internal/repository"
	rabbitmq "github.com/gichohi/go-cqrs-rabbitmq/pkg"
	"log"
)

type CreateProductCommandHandler interface {
	Handle()
}

func (command *CreateProductCommand) Handle() {
	user := &models.Product{ProductID: command.ProductID, Title: command.Title, Description: command.Description, Price: command.Price}

	res,err := repository.CreateProduct(user)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

	exchange := "products"
	queue := "created"

	msgBytes, _ := json.Marshal(user)

	msg := string(msgBytes)

	err = rabbitmq.Publish(msg, queue, exchange)
	if err != nil {
		log.Println("RabbitMQ Error: ", err)
	}
}