package api

import (
	"boilerplate/pkg/logger"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/qri-io/jsonschema"
)

// Response interface for standard API results
type Response struct {
	Result interface{} `json:"result"`
	Error  interface{} `json:"error,omitempty"`
}

// ErrorResponse interface for errors returned via the API
type ErrorResponse struct {
	Code    interface{} `json:"code"`
	Message interface{} `json:"message"`
	Invalid interface{} `json:"invalid,omitempty"`
}

var (
	errInvalidJSON    = errors.New("JSON is invalid")
	errInvalidPayload = errors.New("Payload is invalid")
)

// resultResponseJSON will write the result as json to the http output
func resultResponseJSON(w http.ResponseWriter, status int, result interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	var response Response
	resultClass := fmt.Sprintf("%v", reflect.TypeOf(result))

	if resultClass == "api.ErrorResponse" {
		response = Response{
			Error: result,
		}
	} else {
		response = Response{
			Result: result,
		}
	}

	payload, err := json.Marshal(response)
	if err != nil {
		logger.Error("%+v", err)
	}

	fmt.Fprint(w, string(payload))
}

// resultErrorJSON will format the result as a standard error object and send it for output
func resultErrorJSON(w http.ResponseWriter, status int, message string) {
	errorResponse := ErrorResponse{
		Code:    status,
		Message: message,
	}

	resultResponseJSON(w, status, errorResponse)
}

// resultErrorJSON will format the result as a standard error object and send it for output
func resultSchemaErrorJSON(w http.ResponseWriter, errors []jsonschema.ValError) {
	errorResponse := ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: "Request failed validation",
		Invalid: errors,
	}

	resultResponseJSON(w, http.StatusBadRequest, errorResponse)
}

// validateRequestSchema takes a Schema and the Content to validate against it
func validateRequestSchema(schema string, requestBody []byte) ([]jsonschema.ValError, error) {
	var jsonErrors []jsonschema.ValError
	var schemaBytes = []byte(schema)

	// Make sure the body is valid JSON
	if !isJSON(requestBody) {
		return jsonErrors, errInvalidJSON
	}

	rs := &jsonschema.RootSchema{}
	if err := json.Unmarshal(schemaBytes, rs); err != nil {
		return jsonErrors, err
	}

	var validationErr error
	if jsonErrors, validationErr = rs.ValidateBytes(requestBody); len(jsonErrors) > 0 {
		return jsonErrors, validationErr
	}

	// Valid
	return nil, nil
}

func isJSON(bytes []byte) bool {
	var js map[string]interface{}
	return json.Unmarshal(bytes, &js) == nil
}
