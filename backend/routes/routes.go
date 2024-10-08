package routes

import (
	"cooolJournal/backend/responses"
	"net/http"
)

// This is a test route.
func SaveProfile(w http.ResponseWriter, r *http.Request) {
	body := responses.ResponseBody{
		"saved":  {"true"},
		"id":     {"jer"},
		"userid": {"cat"},
	}
	responses.Response(w, r, &body, http.StatusOK)
}
