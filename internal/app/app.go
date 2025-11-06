package app

import "net/http"

type App struct {
	mux *http.ServeMux
}

func New() *App {
	a := &App{
		mux: http.NewServeMux(),
	}
	initRoutes(a.mux)
	return a
}

func (a *App) Handler() http.Handler {
	return a.mux
}
