package errors

type InternalError struct {
	Code    int
	Message string
}

func (ie *InternalError) Error() string {
	return ie.Message
}

func (ie *InternalError) SetMessage(message string) *InternalError {
	ie.Message = message

	return ie
}

func NewInternalError(code int, message string) *InternalError {
	return &InternalError{
		Code:    code,
		Message: message,
	}
}

var (
	ErrBadRequest   = NewInternalError(400, "Bad request")
	ErrNotFound     = NewInternalError(404, "Resource not found")
	ErrUnauthorized = NewInternalError(401, "Unauthorized access")

	ErrInternalError = NewInternalError(500, "Internal server error")
)
