package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float32
	SKU         string
	CreatedOn   string
	UpdatedOn   string
	DeletedOn   string
}

var ProductList = []*Product{
	&Product{
		ID:          1,
		Name:        "Football",
		Description: "Mikasa tubeless",
		Price:       3000.00,
		SKU:         "ball001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Boots",
		Description: "Green Mercurial Vapor X CR7",
		Price:       5000.00,
		SKU:         "qty1",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
}

type Products []*Product

func GetProducts() Products {
	return ProductList
}

func (p *Products) ToJson(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}
