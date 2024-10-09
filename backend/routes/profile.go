package routes

import (
	"cooolJournal/backend/responses"
	"net/http"
)

// Updates the profile picture for a specific authenticated user.
func UpdateProfilePic(w http.ResponseWriter, r *http.Request) {
	body := responses.ResponseBody{
		"updated": {"true"},
	}
	responses.Respond(w, r, &body, http.StatusOK)
}

// Updates the username of the authenticated user.
func UpdateUsername(w http.ResponseWriter, r *http.Request) {
	body := responses.ResponseBody{
		"updated": {"true"},
	}
	responses.Respond(w, r, &body, http.StatusOK)
}

// Updates the biography of the authenticated user.
func UpdateBio(w http.ResponseWriter, r *http.Request) {
	body := responses.ResponseBody{
		"updated": {"true"},
	}
	responses.Respond(w, r, &body, http.StatusOK)
}
