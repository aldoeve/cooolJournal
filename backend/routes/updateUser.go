package routes

import (
	"cooolJournal/backend/responses"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var UpdatedUsername struct {
	Email           string `json:"email"`
	EnteredUsername string `json:"enteredUsername"`
}

var UpdatedBio struct {
	Email      string `json:"email"`
	EnteredBio string `json:"enteredBio"`
}

// Updates the profile picture for a specific authenticated user.
func UpdateProfilePic(w http.ResponseWriter, r *http.Request) {
	body := responses.ResponseBody{
		"verified": {"false"},
	}
	responses.Respond(w, r, &body, http.StatusOK)
}

// Updates the username of the authenticated user.
func UpdateUsername(w http.ResponseWriter, r *http.Request) {
	//Reads POST request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	//Parses JSON data
	err = json.Unmarshal(body, &UpdatedUsername)
	if err != nil {
		panic(err)
	}

	fmt.Print("opening database...\n")
	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("UPDATE users SET username = ? WHERE email = ?", UpdatedUsername.EnteredUsername, UpdatedUsername.Email)
	if err != nil {
		http.Error(w, "Failed to update username", http.StatusInternalServerError)
		return
	}

	responseBody := responses.ResponseBody{}
	responses.Respond(w, r, &responseBody, http.StatusOK)
}

// Updates the biography of the authenticated user.
func UpdateBio(w http.ResponseWriter, r *http.Request) {
	//Reads POST request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	//Parses JSON data
	err = json.Unmarshal(body, &UpdatedBio)
	if err != nil {
		panic(err)
	}

	fmt.Print("opening database...\n")
	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Updates Bio into db
	_, err = db.Exec("UPDATE users SET bio = ? WHERE email = ?", UpdatedBio.EnteredBio, UpdatedBio.Email)
	if err != nil {
		http.Error(w, "Failed to update bio", http.StatusInternalServerError)
		return
	}

	responseBody := responses.ResponseBody{}
	responses.Respond(w, r, &responseBody, http.StatusOK)
}
