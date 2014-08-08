package main

import (
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	}

	t.Execute(w, nil)
}

func main() {
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/", index)
	http.HandleFunc("/ueditor/go/controller", controller)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
