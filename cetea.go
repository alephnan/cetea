package main

import (
	"html/template"
	"net/http"
	"time"
)

type Welcome struct {
	Name string
	Time string
}

func main() {
	http.HandleFunc("/", mainhandler)
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
