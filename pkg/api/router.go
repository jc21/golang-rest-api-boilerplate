package api

import (
	"github.com/go-chi/chi"
)

// getRouter ...
func getRouter() chi.Router {
	router := applyPublicMiddleware(chi.NewRouter())

	// These are publicly visible routes, not authentication required
	router.NotFound(notFoundHandler)
	router.MethodNotAllowed(notFoundHandler)
	router.Get("/", healthHandler)
	router.Post("/tokens", newTokenHandler)

	// Add all the restricted routes with their own additional middleware
	router.Group(restrictedRoutes)

	return router
}

// These routes require a valid JWT
func restrictedRoutes(router chi.Router) {
	// Be aware, all previously middleware is still on the stack
	router = applyRestrictedMiddleware(router)

	router.Get("/tokens", refreshTokenHandler)
	// Get a user by ID or "me" to get yourself
	router.Get("/users/{userID:(?:[0-9]+|me)}", userHandler)
}
