package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippet Box"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a snippet..."))
}
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display Create Snippet Form.."))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("listening on port :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
