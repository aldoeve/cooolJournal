package routes

import (
	"cooolJournal/backend/responses"
	"net/http"
	"io"
	"fmt"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"path/filepath"
)

var(
	folderPath = "../database/"
	fileName = "database.db"
	filePath = filepath.Join(folderPath, fileName)
	)

var LoginRequest struct {
	EnteredEmail string `json:"enteredEmail"`
	EnteredPassword string `json:"enteredPass"`
}

// Authenticates the user for login purposes.
func LoginUser(w http.ResponseWriter, r *http.Request) {

	//Reads POST request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	//Parses JSON data
	err = json.Unmarshal(body, &LoginRequest)
    if err != nil {
        panic(err)
    }

	fmt.Print("opening database...\n")
    db, err := sql.Open("sqlite3", filePath)
    if err != nil {
        panic(err)
    }
    defer db.Close()

	//Checks if entered email and password is in the database
	var verification bool

	err = db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE email = ? AND password = ?) AS row_exists;",
		LoginRequest.EnteredEmail, 
		LoginRequest.EnteredPassword,
	).Scan(&verification)

	fmt.Println("User exists:",verification)

	//Returns "true" or "false"
	responseBody := responses.ResponseBody{
		"verified": {fmt.Sprintf("%t", verification)},
	}
	responses.Respond(w, r, &responseBody, http.StatusOK)
}
