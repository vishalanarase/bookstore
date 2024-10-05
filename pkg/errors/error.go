package errors

type APIError struct {
	Status  int
	Message string
}

func (e *APIError) Error() string {
	return e.Message
}

func NewAPIError(status int, message string) *APIError {
	return &APIError{
		Status:  status,
		Message: message,
	}
}
