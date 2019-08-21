package api

import (
	"fmt"
	"net/http"

	"boilerplate/pkg/config"
	"boilerplate/pkg/logger"
)

// StartServer creates a http server
func StartServer() {
	logger.Info("server listening on port %v", config.Env.HTTP.Port)

	http.ListenAndServe(fmt.Sprintf(":%v", config.Env.HTTP.Port), getRouter())
}
