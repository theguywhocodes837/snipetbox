package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from snipetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.Header().Set("Cache-Control", "ppublic")
		w.Header().Set("Cache-Control", "max-age=273727448")
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	//w.Write([]byte("Display a specific snippet..."))
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"name": "Alex"}`))
}
