package main

import (
	"fmt"
	"github.com/shii/finance/model"
	"github.com/shii/finance/routes"
	"net/http"
)

func main() {

	cache := model.SetupCache()

	route := routes.SetupRoutes(cache)

	server := &http.Server{
		Addr:    ":8080",
		Handler: route,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Could not listen on %d: %v\n", 8080, err)
	}
	fmt.Println("Server stopped")
}
