package api

import (
	bjwt "boilerplate/pkg/jwt"
	"boilerplate/pkg/logger"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
)

// Here's where you specify your middleware for the router
func applyPublicMiddleware(router chi.Router) chi.Router {

	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(30 * time.Second))

	// Authentication Middleware:
	privateKey, privateKeyParseErr := bjwt.GetPrivateKey()
	if privateKeyParseErr != nil && privateKey == nil {
		logger.Error(fmt.Sprintf("PrivateKeyParseError: %v", privateKeyParseErr))
	}

	publicKey, publicKeyParseErr := bjwt.GetPublicKey()
	if publicKeyParseErr != nil && publicKey == nil {
		logger.Error(fmt.Sprintf("PublicKeyParseError: %v", publicKeyParseErr))
	}

	tokenAuth := jwtauth.New("RS512", privateKey, publicKey)
	router.Use(jwtauth.Verifier(tokenAuth))

	return router
}

// Here's where you specify your restricted middleware for the router
func applyRestrictedMiddleware(router chi.Router) chi.Router {
	router.Use(authenticatorMiddleware)

	return router
}

// authenticatorMiddleware is a authentication middleware to enforce access from the
// jwtauth.Verifier middleware request context values. The Authenticator sends a 401 Unauthorised
// response for any unverified tokens and passes the good ones through.
func authenticatorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, err := jwtauth.FromContext(r.Context())

		if err != nil || token == nil || !token.Valid {
			resultErrorJSON(w, http.StatusUnauthorized, "Unauthorised")
			return
		}

		// Token is authenticated, continue as normal
		next.ServeHTTP(w, r)
	})
}
