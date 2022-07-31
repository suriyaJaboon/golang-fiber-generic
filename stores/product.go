package stores

import (
	"fg/dtos"
	"log"
)

// Product interface public.
type Product interface {
	FindALL() ([]*dtos.Product, error)
	FindByID(id string) (*dtos.Product, error)
	Save(product dtos.Product) error
}

// product struct private.
type product struct {
	// Database client.
	c any
}

// NewProduct argument database client.
func NewProduct(dbc any) Product {
	return &product{c: dbc}
}

// FindALL this function find all return slices and error
func (p *product) FindALL() ([]*dtos.Product, error) {
	return []*dtos.Product{}, nil
}

// FindByID this function find by id return s
func (p *product) FindByID(id string) (*dtos.Product, error) {
	return &dtos.Product{}, nil
}

// Save this function set product to database then return error.
func (p *product) Save(product dtos.Product) error {
	log.Println(product)

	return nil
}
