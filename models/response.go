package models

type ErrorResponse struct {
	Error string `json:"error"`
	StatusCode int `json:"statusCode"`
}
