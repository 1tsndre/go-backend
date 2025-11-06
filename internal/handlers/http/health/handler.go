package health

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/1tsndre/go-backend/internal/models"
)

func NewHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resp := models.HealthResponse{
			Status: "ok",
			Time:   time.Now().UTC().Format(time.RFC3339),
		}
		_ = json.NewEncoder(w).Encode(resp)
	})
}
