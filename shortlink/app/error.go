package app

type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
	UserMessage  string `json:"user_message"`
}
