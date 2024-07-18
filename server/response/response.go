package response

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/RivasCVA/fetchmojiai-service/api"
)

// Writes a JSON response to the given response writer.
//
// The given data is converted to JSON. An error response is written if the marshal fails.
func Write(w http.ResponseWriter, status int, data any) {
	// encode the given data object
	out, err := json.Marshal(data)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "error encoding the data")
		return
	}

	// write the data to the given writer
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		log.Fatal(fmt.Errorf("Write: failed to write a response: %w", err))
	}
}

// Writes a JSON error response to the given response writer.
//
// The response is based on the generated "api.Error" schema.
func WriteError(w http.ResponseWriter, status int, message string) {
	WriteErrorWithCode(w, status, api.ERROR, message)
}

// Writes a JSON error response with a specific code to the given response writer.
//
// The code is based on the generated "api.ErrorCode" enum.
//
// The response is based on the generated "api.Error" schema.
func WriteErrorWithCode(w http.ResponseWriter, status int, code api.ErrorCode, message string) {
	// encode an error object with the given message
	out, err := json.Marshal(api.Error{
		Status:  int64(status),
		Code:    code,
		Message: message,
	})
	if err != nil {
		log.Fatal(fmt.Errorf("WriteError: failed to marshal the error: %w", err))
	}

	// write the data to the given writer
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		log.Fatal(fmt.Errorf("WriteError: failed to write a response: %w", err))
	}
}
