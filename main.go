package main

import (
	"log"
	"net/http"
	"time"

	"github.com/melsonic/load-balancer/server"
	"github.com/melsonic/load-balancer/util"
)

var lbServer http.Server = http.Server{
	Addr:    util.LBAddress,
	Handler: http.HandlerFunc(util.LBHandler),
}

func main() {
	go func() {
		log.Fatal(server.Server1.ListenAndServe())
	}()

	go func() {
		log.Fatal(server.Server2.ListenAndServe())
	}()

	go func() {
		log.Fatal(server.Server3.ListenAndServe())
	}()

	go func() {
		log.Fatal(server.Server4.ListenAndServe())
	}()

	go func() {
		log.Fatal(lbServer.ListenAndServe())
	}()

	// server health check
	go func() {
		var ticker *time.Ticker = time.NewTicker(15 * time.Second)
		for {
			util.CheckHealth()
			<-ticker.C
		}
	}()

	select {}
}
