package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
)

func main() {

	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", templ.Handler(Dev()))

	log.Fatal(http.ListenAndServe("localhost:3333", nil))

}
