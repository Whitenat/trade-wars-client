package main

import (
	"log"
	"net/http"
)

//Home handler function which writes a byte slice
//containing "Welcome to the thunder dome"
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the thunder dome"))
}

func main() {
//Use the http.NewServeMux() function to initialize a new servemux
//then register the home function as the handler for the "/" URL pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

//Use the http.ListenAndServe() function to start a new web server
//We pass in 2 params: the TCP network address to liston on, and the
//servemux we just created. If http.ListenAndServe() returns an error
//that any error return by http.ListenAndServe() is always non-nil
log.Println("Starting server on :4000")
err := http.ListenAndServe(":4000", mux)
log.Fatal(err)
}