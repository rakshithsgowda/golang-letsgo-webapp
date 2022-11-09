package main

import (
	"log"
	"net/http"
	"text/template"
	"time"
)

// handler function
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi there! We are topot."))
}

func displayTime(w http.ResponseWriter, r *http.Request) {
	// 1. Get the time
	localTime := time.Now().Format("3:04:50PM")
	// 2. read in the template file
	ts, _ := template.ParseFiles("./ui/html/display.time.tmpl")
	// 3. substitute from template usisn-> (template Engine)
	ts.Execute(w, localTime)
}

func main() {
	// create a new serveMux
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// URL for Time
	mux.HandleFunc("/time", displayTime)

	// create a file server
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// create URL mapping to static directory
	mux.Handle("/resource/", http.StripPrefix("/resource", fileServer))

	// create a web server
	log.Println("starting server on port :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
