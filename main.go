package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice containing
// "hello from snippetbox" as response body

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Print a log message to say that the server is starting
	log.Print("starting a server on :4000")

	// Use http.listenAndServe() function to start a new web servcer. We pass in
	// two params: TCP network address to listen on (in this case ":4000")
	// and the new servemux we just created. If http.listenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and terminate the
	// program. Not that any error returned by http.ListenAndServe() is always
	// non-nil

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
