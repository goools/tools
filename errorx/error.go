package errorx

type IError interface {
	error
	Code() int
}

type Error struct {
	code    int
	message string
}

func (e *Error) Error() string {
	return e.message
}

func (e *Error) Code() int {
	return e.code
}

func NewError(code int, err error) *Error {
	return &Error{
		code:    code,
		message: err.Error(),
	}
}
