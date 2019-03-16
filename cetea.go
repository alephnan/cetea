package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/tjarratt/babble"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type TemplateModel_Index struct {
	BuildName string
	BuildTime string
}

type AuthorizationStruct struct {
	Code string
}

var (
	buildName  = babble.NewBabbler().Babble()
	buildTime  = time.Now().Format(time.Stamp)
	logger     = log.New(os.Stdout, "[cetea] ", 0)
	colorGreen = string([]byte{27, 91, 57, 55, 59, 51, 50, 59, 49, 109})
	colorReset = string([]byte{27, 91, 48, 109})
)

func main() {
	logger.Printf("Build: %s %s %s - %s \n", colorGreen, buildName, colorReset, buildTime)

	handle_index(TemplateModel_Index{buildName, buildTime})
	http.HandleFunc("/authorization", authorization)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func handle_index(model TemplateModel_Index) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates := template.Must(template.ParseFiles("template/index.html"))
		if err := templates.ExecuteTemplate(w, "index.html", model); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
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
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	fmt.Println(auth.Code)

	file, err := ioutil.ReadFile("./config/client_secret.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	conf, err := google.ConfigFromJSON(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	token, err := conf.Exchange(oauth2.NoContext, auth.Code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	if token == nil {
		http.Error(w, "No token response received", http.StatusForbidden)
	}

	response, err := json.Marshal(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	io.WriteString(w, string(response))
}
