package httpServer

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"

	"superheroe-api/superheroe-golang-api/src/entity"
)

var users = map[string]string{
	"batman": "1",
	"thor":   "2",
}

// For HMAC signing method, the key can be any []byte. It is recommended to generate
// a key using crypto/rand or something equivalent. You need the same key for signing
// and validating.
var jwtKey []byte

// Signup sign up a user
func (h *HttpServer) Signup(w http.ResponseWriter, r *http.Request) {
	var credentials entity.User
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		// w.WriteHeader(http.StatusBadRequest)
		HandleError(w, "missing client credentials", http.StatusBadRequest)
		return
	}

	// Verify if the user exists in our "db"
	pwd, ok := users[strings.ToLower(credentials.Username)]
	if !ok || pwd != credentials.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	exp := time.Now().Add(time.Minute * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": credentials.Username,
		"exp":  exp, // expires in 1 minute
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, _ := token.SignedString(jwtKey)

	json.NewEncoder(w).Encode(entity.Auth{
		TokenType:   "Bearer",
		AccessToken: tokenString,
		ExpiresIn:   strconv.Itoa(int(exp)),
	})
}

// RefreshToken refresh a user token
func (h *HttpServer) RefreshToken(w http.ResponseWriter, r *http.Request) {
	// extract body from http.Request context
	payload := r.Context().Value("username")
	username, ok := payload.(string)
	if !ok {
		HandleError(w, "missing username", http.StatusBadRequest)
		return
	}

	exp := time.Now().Add(time.Minute * 1).Unix()
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": username,
		"exp":  exp, // expires in 1 minute
	})

	tokenString, _ := jwtToken.SignedString(jwtKey)

	json.NewEncoder(w).Encode(entity.Auth{
		TokenType:   "Bearer",
		AccessToken: tokenString,
		ExpiresIn:   strconv.Itoa(int(exp)),
	})
}
