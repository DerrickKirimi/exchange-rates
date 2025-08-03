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
	gh := g.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", h)
	sm.Handle("/goodbye", gh)

	http.ListenAndServe(":9090", sm)
}
