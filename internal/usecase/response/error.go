package response

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewErrorResponse(status int, message string) *ErrorResponse {
	return &ErrorResponse{
		Status:  status,
		Message: message,
	}
}
