package dto

type ValidationError struct {
	Field  string   `json:"field"`
	Errors []string `json:"errors"`
}

type ValidationFailedResponse struct {
	Errors []ValidationError `json:"errors"`
}

type ErrorResponse struct {
	Message string `json:"message" example:"unable to find users using id"`
}
