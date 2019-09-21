package errors

func New(code string, text string) error {
	return &Error{Code: code, Message: text}
}

// errorString is a trivial implementation of error.
type Error struct {
	Code    string
	Message string
}

func (e *Error) Error() string {
	return e.Code + "::" + e.Message
}
