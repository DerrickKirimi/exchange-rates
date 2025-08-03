package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	g "github.com/DerrickKirimi/exchange-rates/handlers"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	h := g.NewHello(l)
	gh := g.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", h)
	sm.Handle("/goodbye", gh)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	//signal.Notify(sigchan, os.Kill)

	sig := <-sigchan
	l.Println("Received terminate, graceful shutdown", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
