package main

import (
	"fmt"
	"net/http"
)

var buildDir string = "dist"
var port string = "8080"

func main() {
	fs := http.FileServer(http.Dir(buildDir))
	http.Handle("/", http.StripPrefix("/", fs))

	fmt.Printf("Serving `%s` on http://localhost:%s\n", buildDir, port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
