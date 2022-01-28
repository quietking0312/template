package merrors

import (
	"fmt"
	"runtime"
)

type mErrors struct {
	pos  string
	code int32
	s    string
}

func New(code int32, text string) error {
	mErr := &mErrors{
		code: code,
		s:    text,
	}
	if pc, file, line, ok := runtime.Caller(1); ok {
		pcName := runtime.FuncForPC(pc).Name()
		mErr.pos = fmt.Sprintf("%s %s:%d(0x%x)", pcName, file, line, pc)
	}
	return mErr
}

func (e *mErrors) Error() string {
	return e.s
}

func (e *mErrors) Code() int32 {
	return e.code
}

func (e *mErrors) Pos() string {
	return e.pos
}
