package api

import (
	bjwt "boilerplate/pkg/jwt"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// tokenPayload is the structure we expect from a incoming login request
type tokenPayload struct {
	Identity string `json:"identity"`
	Secret   string `json:"secret"`
}

const tokenRequestSchema = `{
	"type": "object",
	"additionalProperties": false,
	"required": [
		"identity",
		"secret"
	],
	"properties": {
		"identity": {
			"type": "string",
			"minLength": 1,
			"maxLength": 255
		},
		"secret": {
			"type": "string",
			"minLength": 1,
			"maxLength": 255
		}
	}
}`

// Also known as a Login, requesting a new token with credentials
func newTokenHandler(w http.ResponseWriter, r *http.Request) {
	// Read the bytes from the body
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resultErrorJSON(w, http.StatusInternalServerError, err.Error())
	}

	// Schema Validation:
	jsonErrors, err := validateRequestSchema(tokenRequestSchema, bodyBytes)
	// General validation error
	if err != nil {
		code := http.StatusInternalServerError
		if err == errInvalidJSON {
			code = http.StatusBadRequest
		}
		resultErrorJSON(w, code, err.Error())
		return
	}

	// JSON Schema errors
	if jsonErrors != nil {
		resultSchemaErrorJSON(w, jsonErrors)
		return
	}

	var payload tokenPayload
	err = json.Unmarshal(bodyBytes, &payload)
	if err != nil {
		resultErrorJSON(w, http.StatusBadRequest, errInvalidPayload.Error())
		return
	}

	// TODO: Use your own methods to log someone in and then return a new Token

	if response, err := bjwt.Generate(123456); err != nil {
		resultErrorJSON(w, http.StatusInternalServerError, err.Error())
	} else {
		resultResponseJSON(w, http.StatusOK, response)
	}
}

// Refresh an existing token by given them a new one with the same claims
func refreshTokenHandler(w http.ResponseWriter, r *http.Request) {

	// TODO: Use your own methods to verify an existing user is
	// able to refresh their token and then give them a new one

	if response, err := bjwt.Generate(123456); err != nil {
		resultErrorJSON(w, http.StatusInternalServerError, err.Error())
	} else {
		resultResponseJSON(w, http.StatusOK, response)
	}
}
