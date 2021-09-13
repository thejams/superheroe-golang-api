package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

// Superheroe
type Superhero struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id" validate:"omitempty"`
	Name  string             `bson:"name,omitempty" json:"name" validate:"required" `
	Alias string             `bson:"alias,omitempty" json:"alias" validate:"required"`
}

// Message
type Message struct {
	MSG string `json:"msg"`
}
