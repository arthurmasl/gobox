package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, you've requested: %s\n", r.URL.Path)
		fmt.Println(r.URL.Query())
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/", fs)

	http.ListenAndServe(":8000", nil)
}
