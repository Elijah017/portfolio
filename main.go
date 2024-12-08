package main

import (
	"github.com/Elijah017/portfolio/internal/routes"
)

func main() {
	if err := routes.InitRoutes(); err == nil {
	}
}
