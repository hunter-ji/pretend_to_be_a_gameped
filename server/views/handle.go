package views

import (
	"html/template"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/handle.html")
	t.Execute(w, nil)
}
