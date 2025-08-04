package data

import "time"

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

func GetProducts() []*Product {
	return ProductList
}
