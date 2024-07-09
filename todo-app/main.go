package main

import (
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	ID    int
	Title string
	Done  bool
}

func main() {
	home := func(w http.ResponseWriter, req *http.Request) {
		tmpl := template.Must(template.ParseFiles("./static/index.html"))
		todos := []Todo{
			{ID: 1, Title: "Learn Go templates", Done: false},
			{ID: 2, Title: "Build a todo app", Done: false},
		}
		tmpl.Execute(w, todos)
	}
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
