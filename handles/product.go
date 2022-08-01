package handles

import (
	"fg/dtos"
	"fg/services"
	"fg/stores"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type product struct {
	serviceProduct services.Product
}

func NewProduct(f *fiber.App, dbc any) {
	//storeProduct := stores.NewProduct("this.connect-database-client")
	storeProduct := stores.NewProduct(dbc)
	serviceProduct := services.NewProduct(storeProduct)

	p := product{serviceProduct: serviceProduct}

	g := f.Group("/product")
	g.Get("/:id/:name", NewHandleParamsParser(p.FindByIDWithProduct))
	g.Get("/:id", NewHandleParamsParser(p.FindByID))
	g.Post("/", NewHandleBodyParser(p.Create))
	g.Put("/:id", NewHandleParamsWithBodyParser(p.UpdateByID))
	g.Delete("/:id", NewHandleParamsParser(p.DeleteByID))
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
	res, err := p.serviceProduct.FindByID(req.ID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p *product) FindByIDWithProduct(req dtos.Params) (*dtos.Product, error) {
	fmt.Println(req)
	return &dtos.Product{}, nil
}

func (p *product) Create(req dtos.ProductDto) (*dtos.Ok, error) {
	if err := p.serviceProduct.Create(req); err != nil {
		return nil, err
	}

	return dtos.OK, nil
}

func (p *product) UpdateByID(id dtos.Params, req dtos.ProductDto) (*dtos.Ok, error) {
	fmt.Println(id, req)
	return dtos.OK, nil
}

func (p *product) DeleteByID(req dtos.Params) (*dtos.Ok, error) {
	fmt.Println(req)
	return dtos.OK, nil
}
