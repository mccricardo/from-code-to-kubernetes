// Source: https://gobyexample.com/http-servers

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Server is starting on port :8090...")
	http.ListenAndServe(":8090", nil)
}
