package main

import (
	"io"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Write "Hello, world!" to the response body
	io.WriteString(w, "Hello, world!?!\n")
}

func main() {
	// Set up a /hello resource handler
	http.HandleFunc("/hello", helloHandler)

	// Listen to port 8080 and wait
	//log.Fatal(http.ListenAndServe(":8080", nil))

	// Listen to HTTPS connections on port 8443 and wait
	log.Fatal(http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil))
}
