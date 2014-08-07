package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/ueditor/go/controller", controller)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
