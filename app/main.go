package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DevOps Pipeline Running with Golang!")
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}
