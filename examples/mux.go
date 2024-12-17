package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "you requested book \"%s\" on page %s\n", title, page)
	fmt.Printf("book: %s, page: %s", title, page)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", GetBook).Methods("GET")

	http.ListenAndServe(":8000", r)
}
