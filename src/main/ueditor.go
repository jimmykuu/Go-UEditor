package main

import (
	"net/http"
)

func controller(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ueditor controller"))
}
