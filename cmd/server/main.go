package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Furrity/web-resume/pkg/app"
	"github.com/Furrity/web-resume/pkg/handlers"
	"github.com/Furrity/web-resume/pkg/middleware"
	"github.com/alexflint/go-arg"
	"github.com/gorilla/mux"
)

type AppConfig struct {
	Debug     bool   `arg:"--debug,env:DEBUG" help:"Enable debug mode"`
	Port      int    `arg:"--port,env:PORT" help:"Specify what port to run in"`
	StaticDir string `arg:"--static,env:STATIC_DIR" help:"Specify static path"`
}

func main() {
	var args AppConfig
	arg.MustParse(&args)
	myApp := app.NewApp(
		args.Debug,
	)

	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(args.StaticDir))))

	// middleware
	r.Use(middleware.IncrementStatsMiddleware(myApp.IncrementRequestCount))

	// html router
	r.HandleFunc("/", handlers.ResumeHandler).Methods("GET")

	// api router
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/metrics", handlers.MetricsHandler(myApp)).Methods("GET")

	port := fmt.Sprintf(":%d", args.Port)
	log.Fatal(http.ListenAndServe(port, r))
}
