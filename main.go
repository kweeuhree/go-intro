// import main package
package main

// import logging and net/http
import (
	// formatted I/O functions
	"log"
	"net/http" // provides functionality for building HTTP servers and clients
	//conversions to and from string representations of basic data types
)

// define a home handler function which writes a byte slice containing
// "hello world" as the response body
// -- http.ResponseWriter parameter provides methods for assembling HTTP response
// -- and sending it to the user
// -- *http.Request parameter is a pointer to a struct which holds information about
// -- the current request (like the HTTP method and the URL being requested)
func home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/". If it
	// doesn't, use the http.NotFound() function to send a 404 response to the client.
	// return from the handler. Failing to return the handler would result
	// in "hello world" message being printed as well
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("hello world"))
}

// snippetView handler
func snippetView(w http.ResponseWriter, r *http.Request) {
	// write a byte slice as the response body
	w.Write([]byte("Display a specific snippet"))
}

// add snippetCreate handler
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request is using POST or not.
	if r.Method != "POST" {
		// If it's not, use the w.WriteHeader() method to send a 405 status
		// code and the w.Write() method to write a "Method Not Allowed"
		// response body. return, so that the subsequent code is not executed.
		// Use the Header().Set() method to add an 'Allow: POST' header to the
		// response header map. The first parameter is the header name, and
		// the second parameter is the header value
		// -- WriteHeader can be called once per response
		// -- once status code has been written it can't be changed
		// -- without WriteHeader '200 OK' status will be sent, to customize
		// -- you must call WriteHeader before any call to Write
		w.Header().Set("Allow", "POST")                    // must be set before calling WriteHeader or Write
		w.Header().Set("Content-Type", "application/json") // if unset manually will be set as 'text/plain'
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}
	// write a byte slice as a response body
	w.Write([]byte("Create a new snippet"))
}

// -- use http.Error() shortcut. This is a lightweight helper function
// -- which takes a given message and status code, then calls the
// -- w.WriteHeader() and w.Write() methods behind-the-scenes.
// func snippetCreate(w http.ResponseWriter, r *http.Request) {
// 		if r.Method != "POST" {
// 			w.Header().Set("Allow", "POST")
// 			// Use the http.Error() function to send a 405 status code and
// 			// "Method Not Allowed" string as the response body.
// 			http.Error(w, "Method Not Allowed", 405)
// 			return
// 		}

// 		w.Write([]byte("Create a new snippet..."))
// 	}

// initialize main point of entry
func main() {
	// use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for "/" URL pattern
	// avoid DefaultServeMux as it is a global variable, any package can access it and register a route,
	// which is a security issue
	// -- each time the server receives a new HTTP request it will pass the request on to the servemux and
	// -- the servemux will check the URL path and dispatch the request to the matching handler
	// -- servemux treats the URL pattern "/" like a catch-all, meaning that all requests will be
	// -- handled by provided function
	mux := http.NewServeMux() // locally-scoped servemux

	// Go’s servemux supports two different types of URL patterns: fixed
	// paths and subtree paths. Fixed paths don’t end with a trailing slash,
	// whereas subtree paths do end with a trailing slash.

	mux.HandleFunc("/", home) // subtree path
	// add handlers for snippetView and snippetCreate
	mux.HandleFunc("/snippet/view", snippetView)     // fixed path
	mux.HandleFunc("/snippet/create", snippetCreate) // fixed path

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
