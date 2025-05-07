package app

import (
	"sync"

	"github.com/Furrity/web-resume/internal/database"
)

type App struct {
	debug bool
	stats struct {
		RequestCount int
		Mu           sync.Mutex
	}
	Queries *database.Queries
}

func NewApp(
	debug bool,
	queries *database.Queries,
) *App {
	return &App{
		debug: debug,
		stats: struct {
			RequestCount int
			Mu           sync.Mutex
		}{RequestCount: 0},
		Queries: queries,
	}

}

func (a *App) Debug() bool {
	return a.debug
}

func (a *App) IncrementRequestCount() {
	a.stats.Mu.Lock()
	defer a.stats.Mu.Unlock()
	a.stats.RequestCount++
}

func (a *App) GetRequestCount() int {
	a.stats.Mu.Lock()
	defer a.stats.Mu.Unlock()
	return a.stats.RequestCount
}
