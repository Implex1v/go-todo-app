package routes

type ApiError struct {
	Message string
}

func NewApiError(m string) *ApiError {
	return &ApiError{Message: m}
}

func ApiErrorOf(e error) *ApiError {
	return NewApiError(e.Error())
}
