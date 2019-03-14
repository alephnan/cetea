package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", mainhandler)
	http.ListenAndServe(":8080", nil)
}

func mainhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi!")
}
