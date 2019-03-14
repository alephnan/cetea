package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", mainhandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func mainhandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "pong")
}
