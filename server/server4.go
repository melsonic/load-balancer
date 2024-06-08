package server

import (
	"fmt"
	"net/http"
	"strings"
)

var Server4Address string = "http://localhost:8084"

func Server4Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Server 4\n")
}

var Server4 *http.Server = &http.Server{
	Addr:    fmt.Sprintf(":%s", strings.Split(Server4Address, ":")[2]),
	Handler: http.HandlerFunc(Server4Handler),
}
