// import main package
package main

// import logging and net/http
import (
	"log"
	"net/http" // provides functionality for building HTTP servers and clients
)

// define a home handler function which writes a byte slice containing
// "hello world" as the response body
// -- http.ResponseWriter parameter provides methods for assembling HTTP response
// -- and sending it to the user
// -- *http.Request parameter is a pointer to a struct which holds information about
// -- the current request (like the HTTP method and the URL being requested)
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

// initialize main point of entry
func main() {
	// use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for "/" URL pattern
	// -- each time the server receives a new HTTP request it will pass the request on to the servemux and
	// -- the servemux will check the URL path and dispatch the request to the matching handler
	// -- servemux treats the URL pattern "/" like a catch-all, meaning that all requests will be
	// -- handled by provided function
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// print message to console
	log.Print("Starting server on :4000")
	// use http.ListenAndServe9) function to start a new web server.
	// pass in two parameters: the TCP network address to listen on (in this case ":4000");
	// the servemux that was jsut created.
	// if an error is returned, use log.Fatal() function to log the error message and exit
	err := http.ListenAndServe(":4000", mux) // required format "host:port"
	// with omitting host the server will listen on all computer's avalable network interfaces
	log.Fatal(err) // any error returned by http.ListenAndServe() is always non-nil
}
