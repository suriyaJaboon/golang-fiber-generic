package dtos

type Product struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type ProductDto struct {
	Name string `json:"name"`
}

type ProductParams struct {
	ID string `params:"id"`
}
