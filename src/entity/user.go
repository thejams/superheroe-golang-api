package entity

// User represents a user logged in the api
type User struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Auth is the sign up response with token information
type Auth struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	ExpiresIn    string `json:"expires_in,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
}
