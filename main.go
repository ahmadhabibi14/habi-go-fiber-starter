package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	tmpl, err := template.ParseFiles("./index.html")
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		fmt.Println("Server works!")
	})
	http.ListenAndServe(":8000", nil)
}
