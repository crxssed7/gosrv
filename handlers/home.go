package handlers

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/layouts/application.html", "templates/home.html"))
	tmpl.Execute(w, nil)
}
