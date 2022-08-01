package services

import (
	"fg/dtos"
	"fg/stores"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product interface {
	FindALL() ([]*dtos.Product, error)
	FindByID(id primitive.ObjectID) (*dtos.Product, error)
	Create(dto dtos.ProductDto) error
	Update(id primitive.ObjectID, dto dtos.ProductDto) error
	Delete(id primitive.ObjectID) error
}

type StoreProduct stores.Store[dtos.Product]

type product struct {
	store StoreProduct
}

func NewProduct(store StoreProduct) Product {
	return &product{store: store}
}

func (p *product) FindALL() ([]*dtos.Product, error) {
	return p.store.FindALL()
}

func (p *product) FindByID(id primitive.ObjectID) (*dtos.Product, error) {
	prod, err := p.store.FindByID(id)
	if err != nil {
		return nil, ErrByID(err)
	}

	return prod, nil
}

func (p *product) Create(dto dtos.ProductDto) error {
	var prod = dtos.Product{
		ID:        primitive.NewObjectID(),
		UUID:      uuid.New().String(),
		Name:      dto.Name,
		CreatedAt: time.Now(),
	}

	return p.store.Create(prod)
}

func (p *product) Update(id primitive.ObjectID, dto dtos.ProductDto) error {
	prod, err := p.store.FindByID(id)
	if err != nil {
		return ErrByID(err)
	}

	prod.Name = dto.Name
	prod.UpdatedAt = time.Now()

	if err = p.store.Update(id, *prod); err != nil {
		return err
	}

	return nil
}

func (p *product) Delete(id primitive.ObjectID) error {
	_, err := p.store.FindByID(id)
	if err != nil {
		return ErrByID(err)
	}

	if err = p.store.Delete(id); err != nil {
		return err
	}

	return nil
}
