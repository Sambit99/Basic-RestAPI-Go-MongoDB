package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Sambit99/Basic-RestAPI-Go-MongoDB/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterUserRoutes(r)
	var srv = &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server running on port:8080")

	log.Fatal(srv.ListenAndServe())
}
