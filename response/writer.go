package response

import (
	"encoding/json"
	"net/http"
)

func WriteErrorResponse(
	w http.ResponseWriter,
	errorMessage string,
	errorCode string,
	httpCode int,
) {
	result := ErrorResponse{
		Error: Error{
			Message: errorMessage,
			Code:    errorCode,
		},
	}

	w.WriteHeader(httpCode)
	WriteJsonResponse(w, result)
}

func WriteJsonResponse(w http.ResponseWriter, response interface{}) {
	jsonResponse, err := json.MarshalIndent(response, "", "    ")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
