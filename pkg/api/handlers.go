package api

import (
	"net/http"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	resultErrorJSON(w, http.StatusNotFound, "Not found")
}
