package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func userHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	resultResponseJSON(w, http.StatusOK, fmt.Sprintf("You got user: %v", userID))
}
