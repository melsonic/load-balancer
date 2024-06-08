package server

import (
	"fmt"
	"net/http"
  "strings"
)

var Server1Address string = "http://localhost:8081"

func Server1Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Server 1\n")
}

var Server1 *http.Server = &http.Server{
  Addr:    fmt.Sprintf(":%s", strings.Split(Server1Address, ":")[2]),
	Handler: http.HandlerFunc(Server1Handler),
}
