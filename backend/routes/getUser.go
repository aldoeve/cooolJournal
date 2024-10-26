package routes

import (
	"cooolJournal/backend/responses"
	"net/http"
	"io"
	"fmt"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
)

var UserByEmail struct {
	EnteredEmail string `json:"enteredEmail"`
	EnteredPassword string `json:"enteredPass"`
}


func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	//Reads POST request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	//Parses JSON data
	err = json.Unmarshal(body, &UserByEmail)
    if err != nil {
        panic(err)
    }

	fmt.Print("opening database...\n")
    db, err := sql.Open("sqlite3", filePath)
    if err != nil {
        panic(err)
    }
    defer db.Close()

	var id int;

	//Gets user_id from email and password
	err = db.QueryRow(
		"SELECT user_id from users WHERE email = ? AND password = ?",
		UserByEmail.EnteredEmail, 
		UserByEmail.EnteredPassword,
	).Scan(&id)

	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Database query error", http.StatusInternalServerError)
		return
	}

	fmt.Println("user exists...");
	responseBody := responses.ResponseBody{
		"id": {fmt.Sprintf("%d", id)},
	}
	responses.Respond(w, r, &responseBody, http.StatusOK)

}