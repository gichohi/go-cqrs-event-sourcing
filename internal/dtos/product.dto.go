package dtos

import uuid "github.com/satori/go.uuid"

type ProductDto struct {
	ProductID   uuid.UUID `json:"product_id" validate:"required"`
	Title        string    `json:"title" validate:"required,gte=0,lte=255"`
	Description  string    `json:"description" validate:"required,gte=0,lte=255"`
	Price 		float64	  `json:"price" validate:"required,gte=0"`
}

type ProductCreateResponseDto struct {
	ProductID uuid.UUID `json:"product_id" validate:"required"`
}
