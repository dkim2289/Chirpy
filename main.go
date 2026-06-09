package main

import (
	"log"
	"net/http"
)

func main() {
	// new http.ServeMux
	mux := http.NewServeMux()
	// new http.Server struct
	// - Server Handler : ServeMux
	// - Address : 8080
	server := &http.Server{Handler: mux, Addr: ":8080"}
	// Listen & Serve
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
