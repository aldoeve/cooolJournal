package routes

import (
	"cooolJournal/backend/responses"
	"net/http"
)

// Creates a user within the database.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body := responses.ResponseBody{
		"saved":  {"true"},
		"id":     {"jer"},
		"userid": {"cat"},
	}
	responses.Respond(w, r, &body, http.StatusOK)
}
