package controller

import (
	"html/template"
	"net/http"
)

func GetIndex(rw http.ResponseWriter, req *http.Request) {
	type Page struct {
		Title string
	}
	p := Page{
		Title: "Website",
	}

	t := template.Must(template.ParseFiles("app/controller/index.html"))
	t.Execute(rw, p)
}
