package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("hello world starting server...")
	http.ListenAndServe(":8000", nil)

}
