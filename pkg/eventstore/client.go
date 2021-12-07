package store

import (
	"context"
	"encoding/json"
	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/gichohi/go-cqrs-rabbitmq/internal/dtos"
	"github.com/gofrs/uuid"
	"log"
)

var Db *esdb.Client

type ProductCreateEvent struct {
	Id            string
	payload		  string
}

func InitStore() (*esdb.Client, error) {
	settings, err := esdb.ParseConnectionString("{connectionString}")

	if err != nil {
		panic(err)
	}

	Db, err := esdb.NewClient(settings)

	if err != nil {
		log.Println(err)
	}

	return Db, err
}

func CreateEvent(productDto *dtos.ProductDto) error {
	prod, _ := json.Marshal(productDto)
	productJson := string(prod)
	createEvent := ProductCreateEvent{
		Id:            uuid.Must(uuid.NewV4()).String(),
		payload: productJson,
	}

	data, err := json.Marshal(createEvent)

	if err != nil {
		log.Printf(err.Error())
	}

	eventData := esdb.EventData{
		ContentType: esdb.JsonContentType,
		EventType:   "ProductCreateEvent",
		Data:        data,
	}

	_, err = Db.AppendToStream(context.Background(), "products-stream", esdb.AppendToStreamOptions{}, eventData)

	if err != nil {
		log.Printf(err.Error())
	}
	return err
}


