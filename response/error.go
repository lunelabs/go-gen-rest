package response

type Error struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}
