package http

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string
}

type DataResponse struct {
	Result any
}

func WriteErrorResponse(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	out, _ := json.Marshal(&ErrorResponse{
		Error: message,
	})

	w.Write([]byte(out))
}

func WriteDataResponse(w http.ResponseWriter, d any) {
	w.Header().Set("Content-Type", "application/json")

	out, _ := json.Marshal(&DataResponse{
		Result: d,
	})

	w.Write([]byte(out))
}
