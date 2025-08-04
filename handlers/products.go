package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DerrickKirimi/exchange-rates/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	jp := data.GetProducts()
	data, err := json.Marshal(jp)
	if err != nil {
		http.Error(rw, "Unable to Marshal json", http.StatusInternalServerError)
	}
	rw.Write(data)
}
