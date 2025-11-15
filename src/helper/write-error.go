package helper

import (
	"encoding/json"
	"net/http"
)

func WriteError(w http.ResponseWriter, errMsg string, status int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"error": errMsg,
	})
}
