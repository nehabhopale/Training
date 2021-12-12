package MyError

type MyError struct {
	errMsg string
}

func New(errorMsg string) *MyError {
	return &MyError{
		errMsg: errorMsg,
	}
}
func (e *MyError) Error() string {
	return e.errMsg
}
