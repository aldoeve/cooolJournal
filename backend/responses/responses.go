package responses

import (
	"encoding/json"
	"net/http"
)

// This should be used to store data that is being sent.
type ResponseBody map[string][]string

// This should be called after the body is figured out to package the response and send it.
func Response(w http.ResponseWriter, r *http.Request, body *ResponseBody, statusCode int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(*body)
}
