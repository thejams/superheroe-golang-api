package httpServer

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

//tokenValidatorMiddleware Validate the request object fields
func tokenValidatorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := validateBearerAuthHeader(r.Header.Get("authorization"))
		if token == "" {
			HandleError(w, "missing access token", http.StatusBadRequest)
			return
		}

		// Parse takes the token string and a function for looking up the key. The latter is especially
		// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
		// head of the token to identify which key to use, but the parsed token (head and claims) is provided
		// to the callback, providing flexibility.
		parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			// is the key that signed the token in auth handler
			return jwtKey, nil
		})
		if err != nil {
			// if token expired, it will get catch here
			HandleError(w, "invalid access token", http.StatusBadRequest)
			return
		}

		if _, ok := parsedToken.Claims.(jwt.MapClaims); !ok || !parsedToken.Valid {
			HandleError(w, "invalid access token", http.StatusBadRequest)
			return
		}

		// validate that user exists in our "DB"
		payload, _ := parsedToken.Claims.(jwt.MapClaims)
		usr := payload["user"].(string)
		_, ok := users[strings.ToLower(usr)]
		if !ok {
			HandleError(w, "invalid user", http.StatusForbidden)
		}

		ctx := context.WithValue(r.Context(), "username", strings.ToLower(usr))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// validateBearerAuthHeader validates incoming `r.Header.Get("Authorization")` header
// and returns token otherwise an empty string.
func validateBearerAuthHeader(authHeader string) string {
	if authHeader == "" {
		return ""
	}

	parts := strings.Split(authHeader, "Bearer")
	if len(parts) != 2 {
		return ""
	}

	token := strings.TrimSpace(parts[1])
	if len(token) < 1 {
		return ""
	}

	return token
}
