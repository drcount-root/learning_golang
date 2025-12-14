package product

import (
	"errors"
	"math/rand"
	"time"
)

type Product struct {
	id              int
	name            string
	description     string
	originalPrice   float64
	discountedPrice float64
	createdAt       time.Time
}

func NewProduct(name, description string, originalPrice, discountedPrice float64) (*Product, error) {
	if name == "" || description == "" || originalPrice < 0 || discountedPrice < 0 || discountedPrice > originalPrice {
		return nil, errors.New("invalid input data for product")
	}

	return &Product{
		id:              rand.Intn(1000),
		name:            name,
		description:     description,
		originalPrice:   originalPrice,
		discountedPrice: discountedPrice,
		createdAt:       time.Now(),
	}, nil
}

func (p *Product) UpdateProductDetails(id int, name, description string, originalPrice, discountedPrice float64) error {
	if name == "" || description == "" || originalPrice < 0 || discountedPrice < 0 || discountedPrice > originalPrice {
		return errors.New("invalid input data for product")
	}

	if p.id != id {
		return errors.New("product ID mismatch")
	}

	p.name = name
	p.description = description
	p.originalPrice = originalPrice
	p.discountedPrice = discountedPrice

	return nil
}
