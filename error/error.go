package error

type Error struct {
	StatusCode int
	Response   interface{}
	Err        error
}

func (r *Error) Error() string {
	return r.Err.Error()
}
