package main

import (
	"log"
	"net/http"

	"github.com/melsonic/load-balancer/server"
	"github.com/melsonic/load-balancer/util"
)

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

	var lbServer http.Server = http.Server{
		Addr:    util.LBAddress,
		Handler: http.HandlerFunc(util.LBHandler),
	}

	go func() {
		log.Fatal(lbServer.ListenAndServe())
	}()

	select {}
}
