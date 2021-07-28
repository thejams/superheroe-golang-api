package entity

// Superheroe
type Superhero struct {
	ID    string `json:"id" validate:"omitempty"`
	Name  string `json:"name" validate:"required"`
	Alias string `json:"alias" validate:"required"`
}

// Message
type Message struct {
	MSG string `json:"msg"`
}
