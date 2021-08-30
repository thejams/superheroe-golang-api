package entity

// Superheroe
type Superhero struct {
	ID    string `bson:"id,omitempty" json:"id" validate:"omitempty"`
	Name  string `bson:"name,omitempty" json:"name" validate:"required" `
	Alias string `bson:"alias,omitempty" json:"alias" validate:"required"`
}

// Message
type Message struct {
	MSG string `json:"msg"`
}
