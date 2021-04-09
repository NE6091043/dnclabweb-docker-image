package main

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	mux := httprouter.New()
	mux.GET("/", index)
	mux.ServeFiles("/assets/css/*filepath", http.Dir("assets/css"))
	mux.ServeFiles("/assets/js/*filepath", http.Dir("assets/js"))
	mux.ServeFiles("/assets/img/*filepath", http.Dir("assets/img"))
	mux.ServeFiles("/assets/libs/*filepath", http.Dir("assets/libs"))
	mux.ServeFiles("/assets/fonts/*filepath", http.Dir("assets/fonts"))
	server := http.Server{
		Addr:    ":8888",
		Handler: mux,
	}
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	t.Execute(w, "")
}
