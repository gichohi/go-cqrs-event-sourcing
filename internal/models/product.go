package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Product struct {
ProductID   uuid.UUID `json:"product_id"`
Title       string    `json:"title"`
Description string    `json:"description"`
Price       float64   `json:"price"`
CreatedAt   time.Time `json:"created_at"`
UpdatedAt   time.Time `json:"updated_at"`
}
