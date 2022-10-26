package error

type Error struct {
	StatusCode   int
	ErrorCode    string
	ErrorMessage string
	Err          error
}

func (r *Error) Error() string {
	return r.Err.Error()
}
