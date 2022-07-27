package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	// SEM RENDERIZAÇÃO DE HTML
	// http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Olá Mundo!"))
	// })

	//COM RENDEZIDAÇÃO HTML
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		u := usuario{
			"João",
			"João.pedro@hotmail.com",
		}

		templates.ExecuteTemplate(w, "home.html", u)
	})

	fmt.Println("Listen on port 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
