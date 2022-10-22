package resp

type ErrorResponse struct {
	Error Error `json:"error"`
}
