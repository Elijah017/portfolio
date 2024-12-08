package routes

import (
	"html/template"
	"log"
	"net/http"
)

const tmpl_dir = "templates/"

func InitRoutes() error {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homeHandler)

	return nil
}

func homeHandler(writer http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles(tmpl_dir+"base.html", tmpl_dir+"home.html")
	if err != nil {
		log.Fatal(err.Error())
	}
	tmpl.ExecuteTemplate(writer, "base", nil)
}
