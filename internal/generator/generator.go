package generator

import (
	"github.com/Elijah017/portfolio/internal/generator/router"
)

var genDir string = "./static/html"
var tmplDir string = "./templates"

func Generate() error {
	rtr, err := router.NewRouter("/")
	if err != nil {
		return err
	}

	if err := rtr.AddRoute("/index.html"); err != nil {
		return err
	}

	if err := rtr.GeneratePages(genDir, tmplDir); err != nil {
		return err
	}

	return nil
}
