// every main go file needs this unless its a file that we inmport and are doing nothing on it
package main

// we import packages fmt for printing and net/http
// for the server
import (
	"fmt"
	"net/http"
)

// we create a hello function
// the takes a two arguments that we need if we are making different endpoints
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

// this is second endpoint
// the inside is just some fancy stuff nothing to do with starting a server
func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

// handleFunc serves the endpoints with the function
// listen server needs the error and port number
func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
