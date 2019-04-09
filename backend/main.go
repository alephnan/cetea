package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/namsral/flag"
	"github.com/tjarratt/babble"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type TemplateModel_Index struct {
	BuildName string
	BuildTime string
	ClientId  string
}

type HealthResponse struct {
}

const (
	DEFAULT_PORT = 8080
)

var (
	buildName = babble.NewBabbler().Babble()
	buildTime = time.Now().Format(time.Stamp)

	logger = log.New(os.Stdout, "[mylogger] ", 0)

	// TODO: Indicate variables global with better name
	googleAuth      *oauth2.Config
	idTokenAudience []string

	HEALTH_RESPONSE = HealthResponse{}
)

func main() {
	port := flag.Int("port", DEFAULT_PORT, "port to listen on")
	dev := flag.Bool("dev", false, "port to listen on")
	flag.Parse()

	logger.Printf("Build: %s %s %s - %s \n", COLOR_GREEN, buildName, COLOR_RESET, buildTime)
	logger.Printf("Dev mode %s %t %s", COLOR_CYAN, *dev, COLOR_RESET)

	file, err := ioutil.ReadFile("./config/client_secret.json")
	if err != nil {
		// TODO: Signal bash script and/or Docker host and get them to terminate.
		panic(err)
	}
	// TODO: Move into auth.go file somehow
	googleAuth, err = google.ConfigFromJSON(file)
	if err != nil {
		panic(err)
	}
	idTokenAudience = []string{googleAuth.ClientID}

	server := startServerInBackground(*port, *dev)

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

func startServerInBackground(port int, dev bool) *http.Server {
	logger.Printf("Running on port: %s %s %s ", COLOR_GREEN, strconv.Itoa(port), COLOR_RESET)
	addr := ":" + strconv.Itoa(port)
	srv := &http.Server{Addr: addr}

	if dev {
		logger.Printf("Serving static files")
		fs := http.FileServer(http.Dir("static"))
		http.Handle("/", http.StripPrefix("", fs))
	}

	http.HandleFunc("/api/health", health)
	// TODO: move into /api/auth
	http.HandleFunc("/api/authorization", auth_Authenticate)
	// TODO: deprecate since overlaps with authorization. Retained for dev purposes.
	// http.HandleFunc("/api/auth/login", auth_Login)
	http.HandleFunc("/api/auth/refresh", auth_Refresh)
	http.HandleFunc("/api/auth/test", auth_AuthTest)
	http.HandleFunc("/api/auth/logout", auth_Logout)

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			logger.Panic(err)
			panic(err)
		}
	}()
	return srv
}

func health(w http.ResponseWriter, r *http.Request) {
	// TODO: this should be refactored into middleware / interceptor
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HEALTH_RESPONSE)
}
