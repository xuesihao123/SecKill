package model

type Error struct {
	ErrCode int
	ErrMsg string
	ErrStr string
}

func NewError(code int, msg string , str string) *Error {
	return &Error{
		ErrCode: code,
		ErrMsg :  msg,
		ErrStr : str,
	}
}

func (err *Error)Error()string  {
	return err.ErrStr
}

