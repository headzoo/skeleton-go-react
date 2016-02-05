package app

import (
    "net/http"
    "html/template"
)

func GetIndex(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        Title string
    }
    p := Page{
        Title: "Website",
    }

    t := template.Must(template.ParseFiles("templates/index.html"))
    t.Execute(rw, p)
}