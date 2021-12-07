package commands

import (
	"fmt"
	"github.com/gichohi/go-cqrs-rabbitmq/internal/models"
	"github.com/gichohi/go-cqrs-rabbitmq/internal/repository"
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

}