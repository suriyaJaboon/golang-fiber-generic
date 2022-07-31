package services

import (
	"fg/dtos"
	"fg/stores"
)

type Product interface {
	Create(dto dtos.ProductDto) error
	FindALL() ([]*dtos.Product, error)
	FindByID(id string) (*dtos.Product, error)
}

type product struct {
	storeProduct stores.Product
}

func NewProduct(storeProduct stores.Product) Product {
	return &product{storeProduct: storeProduct}
}

func (p *product) FindALL() ([]*dtos.Product, error) {
	var prods = []*dtos.Product{
		{
			UUID: "uuid-new-string-0",
			Name: "golang",
		}, {
			UUID: "uuid-new-string-1",
			Name: "fiber",
		},
	}

	return prods, nil
}

func (p *product) FindByID(id string) (*dtos.Product, error) {
	var prod = dtos.Product{
		UUID: id,
		Name: "golang",
	}

	return &prod, nil
}

func (p *product) Create(dto dtos.ProductDto) error {
	var prod = dtos.Product{
		UUID: "uuid-new-string",
		Name: dto.Name,
	}

	return p.storeProduct.Save(prod)
}
