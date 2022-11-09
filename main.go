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
	// create a web server
	log.Println("starting server on port :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
