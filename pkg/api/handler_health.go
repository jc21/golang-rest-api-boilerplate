package api

import (
	"boilerplate/pkg/config"
	"net/http"
)

type healthCheckResponse struct {
	Commit  string `json:"commit"`
	Healthy bool   `json:"healthy"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	health := healthCheckResponse{
		Commit:  config.Commit,
		Healthy: true,
	}

	resultResponseJSON(w, http.StatusOK, health)
}
