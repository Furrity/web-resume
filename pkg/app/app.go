package app

import "sync"

type App struct {
	debug bool
	stats struct {
		RequestCount int
		Mu           sync.Mutex
	}
}

func NewApp(
	debug bool,
) *App {
	return &App{
		debug: debug,
		stats: struct {
			RequestCount int
			Mu           sync.Mutex
		}{RequestCount: 0},
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
