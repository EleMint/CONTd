package contracts

import "net/http"

// Response defines the base response contract scheme
type Response struct {
	StatusCode int               `json:"status"`
	StatusText string            `json:"status_text"`
	Errors     map[string]string `json:"errors"`
	Warnings   map[string]string `json:"warnings"`
}

// NewResponse returns a new response object
func NewResponse(statusCode int, errors map[string]string, warnings map[string]string) Response {
	return Response{
		StatusCode: statusCode,
		StatusText: http.StatusText(statusCode),
		Errors:     errors,
		Warnings:   warnings,
	}
}

// RootResponse defines the root response contract scheme
type RootResponse struct {
	Response Response
}
