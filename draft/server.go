package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	fmt.Println("server started")
	http.ListenAndServe(":8000", nil)
}

func hello(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	fmt.Println(query)
	fmt.Fprintln(writer, "hello, world")
}

func headers(writer http.ResponseWriter, request *http.Request) {
	for name, headers := range request.Header {
		for _, header := range headers {
			fmt.Fprintf(writer, "%v: %v\n", name, header)
		}
	}
}
