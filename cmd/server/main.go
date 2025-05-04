package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Furrity/web-resume/internal/database"
	"github.com/Furrity/web-resume/pkg/app"
	"github.com/Furrity/web-resume/pkg/handlers"
	"github.com/Furrity/web-resume/pkg/middleware"
	"github.com/alexflint/go-arg"
	"github.com/gorilla/mux"

	_ "modernc.org/sqlite"
)

type AppConfig struct {
	Debug        bool   `arg:"--debug,env:DEBUG" help:"Enable debug mode"`
	Port         int    `arg:"--port,env:PORT" help:"Specify what port to run in"`
	StaticDir    string `arg:"--static,env:STATIC_DIR" help:"Specify static path"`
	DatabasePath string `arg:"-db" default:"." help:"Path to the database"`
}

func main() {
	var args AppConfig
	arg.MustParse(&args)
	db, err := sql.Open("sqlite", args.DatabasePath)
	if err != nil {
		log.Fatal("Could not open database.", err)
	}

	myApp := app.NewApp(
		args.Debug,
		database.New(db),
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
