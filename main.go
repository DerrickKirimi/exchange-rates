package main

import (
	"log"
	"net/http"
	"os"

	g "github.com/DerrickKirimi/exchange-rates/handlers"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	h := g.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", h)

	http.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye World")
	})

	http.ListenAndServe(":9090", sm)
}
