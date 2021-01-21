package errors

type Error struct {
	ID      string      `json:"id"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
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

func NewWithData(id, message string, data interface{}) *Error {
	return &Error{
		ID:      id,
		Message: message,
		Data:    data,
	}
}
