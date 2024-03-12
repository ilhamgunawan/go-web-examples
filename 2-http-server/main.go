package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Printf("Module: 2-http-server\nServer is running on port %d\n", 3333)
	http.ListenAndServe(":3333", nil)
}