package main

import (
	"log"
	"net/http"
)

// handler function
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi there! We are topot."))
}

func main() {
	// create a new serveMux
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// create a file server
	fileServer := http.FileServer(http.Dir("./static/"))
	// create URL mapping to static directory
	mux.Handle("/resource/", http.StripPrefix("/resource", fileServer))

	// create a web server
	log.Println("starting server on port :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
