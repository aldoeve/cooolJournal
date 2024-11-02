package routes

import (
	"cooolJournal/backend/responses"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/mattn/go-sqlite3"
)

var UserByEmail struct {
	EnteredEmail    string `json:"enteredEmail"`
	EnteredPassword string `json:"enteredPass"`
}

func GetUserProfile(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")
	username := parts[len(parts)-1]

	fmt.Print("opening database...\n")
	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var userExists bool

	err = db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE username = ?) AS row_exists;",
		username,
	).Scan(&userExists)

	if userExists {
		var bio string
		err = db.QueryRow(
			"SELECT bio FROM users WHERE username = ?;",
			username,
		).Scan(&bio)

		responseBody := responses.ResponseBody{
			"username": {username},
			"bio":      {bio},
		}
		responses.Respond(w, r, &responseBody, http.StatusOK)
		return
	}
	responseBody := responses.ResponseBody{
		"username": {""},
		"bio":      {""},
	}
	responses.Respond(w, r, &responseBody, http.StatusNotFound)
}

func GetUsernameFromJWT(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		responseBody := responses.ResponseBody{
			"verified": {"false"},
			"username": {""},
		}
		responses.Respond(w, r, &responseBody, http.StatusUnauthorized)
		return
	}

	tokenString := cookie.Value
	token, err := verifyToken(tokenString)

	if err != nil {
		responseBody := responses.ResponseBody{
			"verified": {"false"},
			"username": {""},
		}
		responses.Respond(w, r, &responseBody, http.StatusUnauthorized)
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		iss, _ := claims.GetIssuer()

		fmt.Print("opening database...\n")
		db, err := sql.Open("sqlite3", filePath)
		if err != nil {
			panic(err)
		}
		defer db.Close()

		var username string

		err = db.QueryRow("SELECT username FROM users WHERE email = ?", iss).Scan(&username)
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}

		responseBody := responses.ResponseBody{
			"verified": {"true"},
			"username": {username},
		}
		responses.Respond(w, r, &responseBody, http.StatusOK)
		return
	}
	responseBody := responses.ResponseBody{
		"verified": {"false"},
		"username": {""},
	}
	responses.Respond(w, r, &responseBody, http.StatusUnauthorized)

}
