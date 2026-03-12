package app

import (
	"crypto/sha256"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/mail"
)

type AuthMessage struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Success bool `json:"success"`
}

var LogInPage = func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		Log_message("Wrong method recieved")
		return
	}
	http.ServeFile(w, r, "./template/login.html")
}

var SignUpPage = func(w http.ResponseWriter, r *http.Request) {
	// Placeholder for handliing post
	if r.Method == "POST" {
		stringifyMessage, errStringifingMessage := io.ReadAll(r.Body)
		if errStringifingMessage != nil {
			Log_message("Unable to stringify message")
			return
		}
		Log_message("New request: " + string(stringifyMessage))
		// fourOhFour(w, r)
		json.NewEncoder(w).Encode(
			AuthResponse{
				Success: true,
			},
		)
		return
	}

	if r.Method != "GET" {

		return
	}

	http.ServeFile(w, r, "./template/signup.html")

}

var AuthPage = func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fourOhFour(w, r)
		return
	}

	log.Println("Auth requested")

	// Confirm user details
	// Confirm email address
	var userAttempt AuthMessage
	errDecodingAuthMessage := json.NewDecoder(r.Body).Decode(&userAttempt)
	if errDecodingAuthMessage != nil {
		panic(errDecodingAuthMessage)
	}

	_, errParsingEmailAddress := mail.ParseAddress(userAttempt.Email)
	if errParsingEmailAddress != nil {
		Log_message(errDecodingAuthMessage.Error())
		json.NewEncoder(w).Encode(AuthResponse{
			Success: false,
		})
	}

	Log_message("Email address: " + userAttempt.Email + " is valid")

	// Confirm password
	sha256.Sum256([]byte(userAttempt.Password))

	// json.NewEncoder(w).Encode(AuthResponse{
	// 	Success: true,
	// })

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"redirect": "/shorten/"}`))
}
