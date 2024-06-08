package server

import (
	"fmt"
	"net/http"
	"strings"
)

var Server2Address string = "http://localhost:8082"

func Server2Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Server 2\n")
}

var Server2 *http.Server = &http.Server{
	Addr:    fmt.Sprintf(":%s", strings.Split(Server2Address, ":")[2]),
	Handler: http.HandlerFunc(Server2Handler),
}
