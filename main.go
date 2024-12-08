package main

import (
	"log"
	"net/http"

	"github.com/Elijah017/portfolio/internal/routes"
)

func main() {
	if err := routes.InitRoutes(); err == nil {
	}

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
