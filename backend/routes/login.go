package routes

import (
	"cooolJournal/backend/responses"
	"net/http"
)

// Populate this function.
func LoginUser(w http.ResponseWriter, r *http.Request) {
	body := responses.ResponseBody{
		"verfied": {"true"},
	}
	responses.Response(w, r, &body, http.StatusOK)
}
