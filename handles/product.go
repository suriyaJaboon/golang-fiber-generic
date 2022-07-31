package handles

import (
	"fg/dtos"
	"fg/services"
	"fg/stores"
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
	g.Get("/:id", NewHandleParamsParser(p.FindByID))
	g.Post("/", NewHandleBodyParser(p.Create))
	//g.Post("/", NewHandleBodyParser[dtos.ProductDto, *dtos.Ok](p.Create))
}

func (p *product) FindALL() ([]*dtos.Product, error) {
	products, err := p.serviceProduct.FindALL()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *product) FindByID(req dtos.ProductParams) (*dtos.Product, error) {
	res, err := p.serviceProduct.FindByID(req.ID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p *product) Create(req dtos.ProductDto) (*dtos.Ok, error) {
	if err := p.serviceProduct.Create(req); err != nil {
		return nil, err
	}

	return dtos.OK, nil
}
