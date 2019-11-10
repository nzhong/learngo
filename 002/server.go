package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func logBasicConfigs() string {
	serverPort := getEnv("GO_SERVER_PORT", "8080")
	log.Printf("Server will run on: %s\n", serverPort)
	return serverPort
}

func handleRequest1(res http.ResponseWriter, req *http.Request) {
	reqMethodStr := "echo1 Method: " + req.Method
	//log.Printf("echo1 Method: %v", req.Method)

	res.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
	res.WriteHeader(http.StatusOK)
	io.WriteString(res, reqMethodStr+"\n")

	for name, headers := range req.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			io.WriteString(res, fmt.Sprintf("  %v - %v", name, h)+"\n")
		}
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)
	io.WriteString(res, buf.String()+"\n")
}

func handleRequest2(res http.ResponseWriter, req *http.Request) {
	reqMethodStr := "echo2 Method: " + req.Method
	//log.Printf("echo2 Method: %v", req.Method)

	res.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
	res.WriteHeader(http.StatusOK)
	io.WriteString(res, reqMethodStr+"\n")

	for name, headers := range req.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			io.WriteString(res, fmt.Sprintf("  %v - %v", name, h)+"\n")
		}
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)
	io.WriteString(res, buf.String()+"\n")
}

func main() {
	serverPort := logBasicConfigs()

	http.HandleFunc("/echo1", handleRequest1)
	http.HandleFunc("/echo2", handleRequest2)
	if err := http.ListenAndServe(":"+serverPort, nil); err != nil {
		log.Printf("Error: %v", err)
		panic(err)
	}
}
