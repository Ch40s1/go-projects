package main

import "net/http"

func main() {
	// create new muliplexer
	mux := http.NewServeMux()
	// register the routes and handlers
	mux.Handle("/", &homeHandler{})
	//run server
	http.ListenAndServe(":8080", mux)
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}
