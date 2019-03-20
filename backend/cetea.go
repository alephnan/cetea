package main

import (
	"context"
	"encoding/json"
	"flag"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	verifier "github.com/alephnan/google-auth-id-token-verifier"
	"github.com/tjarratt/babble"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	crm "google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/option"
)

type TemplateModel_Index struct {
	BuildName string
	BuildTime string
	ClientId  string
}

type AuthorizationStruct struct {
	Code     string
	Id_Token string
}

// TODO: populate more fields
// TODO: Support orgs and folders.
type AuthorizationResponse struct {
	Projects []string `json:"projects"`
}

var (
	buildName   = babble.NewBabbler().Babble()
	buildTime   = time.Now().Format(time.Stamp)
	defaultPort = 8080

	logger     = log.New(os.Stdout, "[cetea] ", 0)
	colorGreen = string([]byte{27, 91, 57, 55, 59, 51, 50, 59, 49, 109})
	colorReset = string([]byte{27, 91, 48, 109})

	googleAuth            *oauth2.Config
	googleIdTokenVerifier = verifier.Verifier{}
	idTokenAudience       []string
)

func main() {
	logger.Printf("Build: %s %s %s - %s \n", colorGreen, buildName, colorReset, buildTime)

	port := flag.Int("port", defaultPort, "port to listen on")
	flag.Parse()

	file, err := ioutil.ReadFile("./config/client_secret.json")
	if err != nil {
		// TODO: Signal bash script and/or Docker host and get them to terminate.
		panic(err)
	}
	googleAuth, err = google.ConfigFromJSON(file)
	if err != nil {
		panic(err)
	}
	idTokenAudience = []string{googleAuth.ClientID}

	server := startServerInBackground(*port)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	sig := <-stop
	logger.Println()
	logger.Println("Received signal", sig.String())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err := server.Shutdown(ctx); err != nil {
		cancel()
		logger.Println("Error shutting down.")
	} else {
		logger.Println("Graceful shutdown.")
	}
	logger.Println("Exiting.")
}

func startServerInBackground(port int) *http.Server {
	logger.Printf("Running on port: %s %s %s ", colorGreen, strconv.Itoa(port), colorReset)
	addr := ":" + strconv.Itoa(port)
	srv := &http.Server{Addr: addr}
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	handle_index(TemplateModel_Index{buildName, buildTime, googleAuth.ClientID})
	http.HandleFunc("/authorization", authorization)
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			logger.Panic(err)
			panic(err)
		}
	}()
	return srv
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

	_, err = verifyIdToken(auth.Id_Token)
	if err != nil {
		http.Error(w, "Cannot verify id_token JWT", http.StatusForbidden)
		return
	}

	token, err := googleAuth.Exchange(oauth2.NoContext, auth.Code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	if token == nil {
		http.Error(w, "No token response received", http.StatusForbidden)
	}

	ctx := context.Background()
	crmService, err := crm.NewService(ctx, option.WithTokenSource(googleAuth.TokenSource(ctx, token)))
	projectsResponse, err := crmService.Projects.List().Do()
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	projects := projectsResponse.Projects
	// TODO: handle non 200 HTTP responses?
	// TODO: handle empty project list
	var projectNames = make([]string, len(projects))
	for i := 0; i < len(projects); i++ {
		projectNames[i] = projects[i].Name
	}
	responseStruct := AuthorizationResponse{Projects: projectNames}
	response, err := json.Marshal(responseStruct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	io.WriteString(w, string(response))
}

func verifyIdToken(idToken string) (*verifier.ClaimSet, error) {
	logger.Printf("Verifying id_token: " + idToken)
	err := googleIdTokenVerifier.VerifyIDToken(idToken, idTokenAudience)
	if err != nil {
		logger.Printf("Error verifying id_token.")
		return nil, err
	}
	claims, err := verifier.Decode(idToken)
	if err != nil {
		logger.Print("Error decoding id_token.")
	}
	return claims, err
}
