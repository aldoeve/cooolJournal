package routes

import (
	"cooolJournal/backend/responses"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"time"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/mattn/go-sqlite3"
)

var secretKey = []byte("secret-key")

var (
	folderPath = "../database/"
	fileName   = "database.db"
	filePath   = filepath.Join(folderPath, fileName)
)

var LoginRequest struct {
	EnteredEmail    string `json:"enteredEmail"`
	EnteredPassword string `json:"enteredPass"`
}

var UserCreation struct {
	EnteredEmail    string `json:"enteredEmail"`
	EnteredPassword string `json:"enteredPass"`
}

// Creates a user within the database.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	//Reads POST request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	//Parses JSON data
	err = json.Unmarshal(body, &UserCreation)
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
	var exists bool

	_ = db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE email = ? AND password = ?) AS row_exists;",
		UserCreation.EnteredEmail,
		UserCreation.EnteredPassword,
	).Scan(&exists)

	//User doesn't exists
	if !exists {

		fmt.Println("creating user...")

		//Adds User to database
		_, err = db.Exec(
			"INSERT INTO users (email, username, password, terms, bio) VALUES (?, ?, ?, ?, ?)",
			UserCreation.EnteredEmail, "User", UserCreation.EnteredPassword, time.Now(), "")

		//Creates JWT token
		nowTime := time.Now()
		expirationTime := nowTime.Add(12 * time.Hour)
		tokenString, err := CreateToken(UserCreation.EnteredEmail, expirationTime)
		if err != nil {
			fmt.Errorf("No username found")
		}

		//Set up a Cookie
		http.SetCookie(w,
			&http.Cookie{
				Name:     "token",
				Value:    tokenString,
				Expires:  expirationTime,
				Path:     "/",
				HttpOnly: true,
			})

		responseBody := responses.ResponseBody{
			"userExists": {fmt.Sprintf("%t", exists)},
			"error":      {""},
			"token":      {tokenString},
		}
		responses.Respond(w, r, &responseBody, http.StatusOK)
	} else { //Does exist
		fmt.Println("user already exists...")
		responseBody := responses.ResponseBody{
			"userExists": {fmt.Sprintf("%t", exists)},
			"error":      {"User already exists"},
		}
		responses.Respond(w, r, &responseBody, http.StatusConflict)
	}

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
	var userExists bool

	err = db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE email = ? AND password = ?) AS row_exists;",
		LoginRequest.EnteredEmail,
		LoginRequest.EnteredPassword,
	).Scan(&userExists)

	//The correct email and password were entered
	if userExists {
		//Create JWT token
		nowTime := time.Now()
		expirationTime := nowTime.Add(12 * time.Hour)
		tokenString, err := CreateToken(UserCreation.EnteredEmail, expirationTime)

		if err != nil {
			fmt.Errorf("No username found")
		}

		//Set up a Cookie
		http.SetCookie(w,
			&http.Cookie{
				Name:     "token",
				Value:    tokenString,
				Expires:  expirationTime,
				Path:     "/",
				HttpOnly: true,
			})

		fmt.Println("user exists...")
		responseBody := responses.ResponseBody{
			"userExists": {fmt.Sprintf("%t", userExists)},
			"error":      {""},
			"token":      {tokenString},
		}
		responses.Respond(w, r, &responseBody, http.StatusOK)
	} else {
		fmt.Println("user doesn't exists...")
		responseBody := responses.ResponseBody{
			"userExists": {fmt.Sprintf("%t", userExists)},
			"error":      {"User doesn't exist"},
		}
		responses.Respond(w, r, &responseBody, http.StatusConflict)
	}

}

// Creates a JWT Token
func CreateToken(email string, expirationTime time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    email,
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		responseBody := responses.ResponseBody{
			"verified": {"false"},
			"error":    {"couldn't find cookie"},
		}
		responses.Respond(w, r, &responseBody, http.StatusUnauthorized)
		return
	}

	tokenString := cookie.Value
	_, err = verifyToken(tokenString)

	if err != nil {
		responseBody := responses.ResponseBody{
			"verified": {"false"},
			"error":    {"token not verified"},
		}
		responses.Respond(w, r, &responseBody, http.StatusUnauthorized)
		return
	}

	responseBody := responses.ResponseBody{
		"verified": {"true"},
		"error":    {""},
	}
	responses.Respond(w, r, &responseBody, http.StatusOK)

}

// Returns email based on jwt inside cookie
func RetrieveUser(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("token")
	if err != nil {
		responseBody := responses.ResponseBody{
			"verified": {"false"},
			"emai":     {""},
		}
		responses.Respond(w, r, &responseBody, http.StatusUnauthorized)
		return
	}

	tokenString := cookie.Value
	token, err := verifyToken(tokenString)

	if err != nil {
		responseBody := responses.ResponseBody{
			"verified": {"false"},
			"emai":     {""},
		}
		responses.Respond(w, r, &responseBody, http.StatusUnauthorized)
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		iss, _ := claims.GetIssuer()
		responseBody := responses.ResponseBody{
			"verified": {"true"},
			"email":    {iss},
		}
		responses.Respond(w, r, &responseBody, http.StatusOK)
		return
	}
	responseBody := responses.ResponseBody{
		"verified": {"false"},
		"email":    {""},
	}
	responses.Respond(w, r, &responseBody, http.StatusUnauthorized)

}

// Verifies that jwt is valid
func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return token, err
	}

	if !token.Valid {
		return token, fmt.Errorf("invalid token")
	}

	return token, nil
}
