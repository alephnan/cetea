package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"time"
)

type Welcome struct {
	Name string
	Time string
}

type AuthorizationStruct struct {
	Code string
}

func main() {
	http.HandleFunc("/", mainhandler)
	http.HandleFunc("/authorization", authorization)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func mainhandler(w http.ResponseWriter, r *http.Request) {
	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}
	if name := r.FormValue("name"); name != "" {
		welcome.Name = name
	}
	templates := template.Must(template.ParseFiles("template/index.html"))
	if err := templates.ExecuteTemplate(w, "index.html", welcome); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// io.WriteString(w, "pong")
}

func authorization(w http.ResponseWriter, r *http.Request) {
	// https://stackoverflow.com/questions/17478731/whats-the-point-of-the-x-requested-with-header
	if xRequestedWithHeader := r.Header.Get("X-Requested-With"); xRequestedWithHeader != "XMLHttpRequest" {
		http.Error(w, "Untrusted request", http.StatusForbidden)
		return
	}
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	var auth AuthorizationStruct
	err := json.NewDecoder(r.Body).Decode(&auth)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println(auth.Code)

	data, err := json.Marshal(auth)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	io.WriteString(w, string(data))
}
