package response

type ErrorResponse struct {
	Error Error `json:"error"`
}
