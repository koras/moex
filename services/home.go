package services

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	tmpl, _ := template.ParseFiles("html/home.html")
	tmpl.Execute(w, "")
}
