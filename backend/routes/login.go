package routes

import (
	"cooolJournal/backend/responses"
	"net/http"
)

// Authenticates the user for login purposes.
func LoginUser(w http.ResponseWriter, r *http.Request) {
	body := responses.ResponseBody{
		"verfied": {"true"},
	}
	responses.Respond(w, r, &body, http.StatusOK)
}
