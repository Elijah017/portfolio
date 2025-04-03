package main

import (
	"fmt"
	"net/http"
)

var distDir string = "dist"
var port string = "8080"

func main() {
	fs := http.FileServer(http.Dir(distDir))
	http.Handle("/", http.StripPrefix("/", fs))

	fmt.Printf("Serving `%s` on http://localhost:%s\n", distDir, port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
