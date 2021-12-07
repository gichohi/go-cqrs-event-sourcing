package api

import (
	"encoding/json"
	"fmt"
	"github.com/gichohi/go-cqrs-rabbitmq/internal/commands"
	"github.com/gichohi/go-cqrs-rabbitmq/internal/dtos"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)


func NewHandler() http.Handler {
	router := chi.NewRouter()
	router.Route("/products", routes)
	return router
}

func routes(router chi.Router) {
	router.Post("/", CreateProduct)
	router.Get("/", Test)

}

func Test(w http.ResponseWriter, r *http.Request){
	_, err := w.Write([]byte("<h1>Welcome to my server</h1>"))
	if err != nil {
		log.Fatal(err)
	}

}

func CreateProduct(w http.ResponseWriter, r *http.Request){
	var productDto dtos.ProductDto

	err := json.NewDecoder(r.Body).Decode(&productDto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	command := commands.CreateProductCommand{ProductID: productDto.ProductID, Title: productDto.Title, Description: productDto.Description, Price: productDto.Price}
	command.Handle()

	if err != nil {
		fmt.Println("Error: ", err)
	}
	msg, err := json.Marshal(dtos.ProductCreateResponseDto{ProductID: productDto.ProductID})
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(msg)
	if err != nil {
		log.Fatalf("Write Error: %s", err)
	}
}
