package app

import (
	"net/http"

	"github.com/1tsndre/go-backend/internal/handler/http/health"
)

func initRoutes(mux *http.ServeMux) {
	mux.Handle("/health", health.NewHandler())
}
