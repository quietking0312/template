package merrors

type mErrors struct {
	code int32
	s    string
}

func New(code int32, text string) error {
	return &mErrors{
		code: code,
		s:    text,
	}
}

func (e *mErrors) Error() string {
	return e.s
}

func (e *mErrors) Code() int32 {
	return e.code
}
