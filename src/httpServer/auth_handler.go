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
	"spawn":  "3",
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
	if pwd, ok := users[strings.ToLower(credentials.Username)]; !ok || pwd != credentials.Password {
		HandleError(w, "invalid client credentials", http.StatusUnauthorized)
		return
	}

	exp := time.Now().Add(time.Minute * 1).Unix()
	p := make(map[string]interface{})
	p["user"] = credentials.Username
	p["exp"] = exp
	p["type"] = "access"

	token := signToken(p)
	// Sign and get the complete encoded token as a string using the secret
	accessTokenString, _ := token.SignedString(jwtKey)

	exp = time.Now().Add(time.Minute * 5).Unix()
	p["user"] = credentials.Username
	p["exp"] = exp
	p["type"] = "refresh"

	token = signToken(p)
	refreshTokenString, _ := token.SignedString(jwtKey)

	json.NewEncoder(w).Encode(entity.Auth{
		TokenType:    "Bearer",
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
		ExpiresIn:    strconv.Itoa(int(exp)),
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

	p := make(map[string]interface{}) // jwt.MapClaims es = type MapClaims map[string]interface{}
	p["user"] = username
	p["exp"] = exp

	jwtToken := signToken(p)
	tokenString, _ := jwtToken.SignedString(jwtKey)

	json.NewEncoder(w).Encode(entity.Auth{
		TokenType:   "Bearer",
		AccessToken: tokenString,
		ExpiresIn:   strconv.Itoa(int(exp)),
	})
}

func signToken(payload jwt.MapClaims) *jwt.Token {
	/*
		jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user": credentials.Username,
			"exp":  exp, // expires in 1 minute
		})
	*/
	return jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
}
