package handles

import (
	"fg/dtos"
	"fg/services"
	"fg/stores"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type product struct {
	serviceProduct services.Product
}

func NewProduct(f fiber.Router, mgc *stores.MONGOClient) {
	storeProduct := stores.NewStore[dtos.Product](mgc)
	serviceProduct := services.NewProduct(storeProduct)

	prod := product{serviceProduct: serviceProduct}

	g := f.Group("/product")
	g.Get("", NewHandleResponse(prod.FindALL))
	g.Get("/:id/:name", NewHandleParamsParser(prod.FindByIDWithProduct))
	g.Get("/:id", NewHandleParamsParser(prod.FindByID))
	g.Post("/", NewHandleBodyParser(prod.Create))
	g.Put("/:id", NewHandleParamsWithBodyParser(prod.UpdateByID))
	g.Delete("/:id", NewHandleParamsParser(prod.DeleteByID))
	//g.Post("/", NewHandleBodyParser[dtos.ProductDto, *dtos.Ok](p.Create))
}

func (p *product) FindALL() ([]*dtos.Product, error) {
	products, err := p.serviceProduct.FindALL()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *product) FindByID(req dtos.Params) (*dtos.Product, error) {
	idx, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		return nil, ErrInvalidID
	}

	res, err := p.serviceProduct.FindByID(idx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p *product) FindByIDWithProduct(req dtos.Params) (*dtos.Product, error) {
	return &dtos.Product{}, nil
}

func (p *product) Create(req dtos.ProductDto) (*dtos.Ok, error) {
	if err := p.serviceProduct.Create(req); err != nil {
		return nil, err
	}

	return dtos.OK, nil
}

func (p *product) UpdateByID(id dtos.Params, req dtos.ProductDto) (*dtos.Ok, error) {
	idx, err := primitive.ObjectIDFromHex(id.ID)
	if err != nil {
		return nil, ErrInvalidID
	}

	if err = p.serviceProduct.Update(idx, req); err != nil {
		return nil, err
	}

	return dtos.OK, nil
}

func (p *product) DeleteByID(req dtos.Params) (*dtos.Ok, error) {
	idx, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		return nil, ErrInvalidID
	}

	if err = p.serviceProduct.Delete(idx); err != nil {
		return nil, err
	}

	return dtos.OK, nil
}
