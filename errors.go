package errors

type Error struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func New(id, message string) *Error {
	return &Error{
		ID:      id,
		Message: message,
	}
}
