package main

import (
	"flag"
	"fmt"

	"github.com/Elijah017/portfolio/internal/generator"
)

var distDir = "dist"

func main() {
	generateFlag := flag.Bool("no-gen", false, "whether to generate the pages")
	buildFlag := flag.Bool("no-build", false, "whether to build distributables")
	deployFlag := flag.Bool("no-deploy", false, "whether to deploy to cloudflare")
	flag.Parse()

	if !*generateFlag {
		fmt.Println("generating...")
		generator.Generate()
	}

	if !*buildFlag {
		fmt.Println("building...")
	}

	if !*deployFlag {
		fmt.Println("deploying...")
	}
}
