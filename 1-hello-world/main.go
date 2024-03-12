package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HttpHandler)
	fmt.Printf("Server is running on port %d\n", 3333)
	http.ListenAndServe(":3333", nil)
}

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you are requesting %s\n", r.URL.Path)
}