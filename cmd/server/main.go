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
	"github.com/go-chi/chi/v5"
	cmiddleware "github.com/go-chi/chi/v5/middleware"

	_ "modernc.org/sqlite"
)

type AppConfig struct {
	Debug        bool   `arg:"--debug,env:DEBUG" help:"Enable debug mode"`
	Port         int    `arg:"--port,env:PORT" help:"Specify what port to run in"`
	StaticDir    string `arg:"--static,env:STATIC_DIR" help:"Specify static path"`
	DatabasePath string `arg:"--db" default:"." help:"Path to the database"`
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

	r := chi.NewRouter()
	r.Use(cmiddleware.Logger)
	r.Use(middleware.IncrementStatsMiddleware(myApp.IncrementRequestCount))

	r.Route("/static", func(r chi.Router) {
		fs := http.StripPrefix("/static/", http.FileServer(http.Dir(args.StaticDir)))
		r.Handle("/*", fs)
	})

	// middleware

	// html router
	r.Get("/", handlers.ResumeHandler)

	// api router

	api := chi.NewRouter()
	r.Mount("/api", api)
	api.Get("/metrics", handlers.MetricsHandler(myApp))

	port := fmt.Sprintf(":%d", args.Port)
	log.Fatal(http.ListenAndServe(port, r))
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
}
