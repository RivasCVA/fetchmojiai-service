// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package api

// Defines values for ErrorCode.
const (
	ERROR ErrorCode = "ERROR"
)

// Error A generic error
type Error struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Status  int64     `json:"status"`
}

// ErrorCode defines model for ErrorCode.
type ErrorCode string

// BadRequest A generic error
type BadRequest = Error

// InternalServerError A generic error
type InternalServerError = Error

// ImagineJSONBody defines parameters for Imagine.
type ImagineJSONBody struct {
	// Prompt A prompt to imagine
	Prompt string `json:"prompt"`
}

// ImagineJSONRequestBody defines body for Imagine for application/json ContentType.
type ImagineJSONRequestBody ImagineJSONBody
