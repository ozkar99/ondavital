package main

import (
	"fmt"
	"github.com/ozkar99/ondavital"
	"github.com/VividCortex/robustly"
	"html/template"
	"net/http"
)

func main() {

	robustly.Run(func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			t, err := template.ParseFiles("index.html")
			if err != nil {
				fmt.Println("Error: ", err)
			}

			q := r.FormValue("q")
			result, _ := ondavital.Search(q)
			t.Execute(w, result)
		})

		http.ListenAndServe(":8080", nil)
	}, nil)
}
