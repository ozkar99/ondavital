package main

import (
	"fmt"
	"github.com/ozkar99/ondavital"
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("index.html")
		if err != nil {
			fmt.Println("Error: ", err)
		}

		q := r.FormValue("q")
		result, err := ondavital.Search(q)
		if err != nil {
			result = "Sorry, Error..."
		}

		t.Execute(w, result)
	})

	http.ListenAndServe(":8080", nil)
}
