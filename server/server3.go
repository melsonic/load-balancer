package server

import (
	"fmt"
	"net/http"
	"strings"
)

var Server3Address string = "http://localhost:8083"

func Server3Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Server 3\n")
}

var Server3 *http.Server = &http.Server{
	Addr:    fmt.Sprintf(":%s", strings.Split(Server3Address, ":")[2]),
	Handler: http.HandlerFunc(Server3Handler),
}
