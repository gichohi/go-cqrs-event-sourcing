package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gichohi/go-cqrs-rabbitmq/internal/db"
	"github.com/gichohi/go-cqrs-rabbitmq/internal/models"
	"log"
)

func  CreateProduct(product *models.Product) (sql.Result, error) {

	sqlStatement := `
	INSERT INTO products (product_id, title, description, price, created_at, updated_at)
	VALUES ($1, $2, $3, $4, now(), now())
	`
	statement, err := db.Db.Prepare(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}

	res, err := statement.Exec(product.ProductID, product.Title, product.Description, product.Price)

	if err != nil {
		fmt.Println(err)
	}

	return res, err
}

func Test(product *models.Product) string {
	json, _ := json.Marshal(product)
	return string(json)
}