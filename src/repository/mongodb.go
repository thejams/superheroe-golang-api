//Package repository provides all the methods to persit data in a slice
package repository

import (
	"context"
	"fmt"
	"os"
	"strings"
	"superheroe-api/superheroe-golang-api/src/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// repository main struct for mongodb connection
type repository struct {
	db *mongo.Database
}

// getDBConnection provides all the string connections for the database
func getDBConnection() (string, string, string, string) {
	usr := os.Getenv("MONGO_USER")
	if len(strings.TrimSpace(usr)) == 0 {
		usr = "test"
	}
	pwd := os.Getenv("MONGO_PWD")
	if len(strings.TrimSpace(pwd)) == 0 {
		pwd = "12345"
	}
	host := os.Getenv("MONGO_HOST")
	if len(strings.TrimSpace(host)) == 0 {
		host = "localhost"
	}
	port := os.Getenv("MONGO_PORT")
	if len(strings.TrimSpace(port)) == 0 {
		port = "27017"
	}
	return usr, pwd, host, port
}

// NewMongoConnection provides a new mongodb connection
func NewMongoConnection(ctx context.Context) (Repository, *mongo.Client) {
	usr, pwd, host, port := getDBConnection()
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", usr, pwd, host, port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	// defer client.Disconnect(ctx)

	database := client.Database("user-test")
	return &repository{
		db: database,
	}, client
	/* podcastsCollection := database.Collection("podcasts")
	episodesCollection := database.Collection("episodes") */
}

func (r *repository) GetSuperheroes(ctx context.Context) []entity.Superhero {
	var superheroes []entity.Superhero
	collection := r.db.Collection("superheroe")
	filter := bson.M{}
	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &superheroes); err != nil {
		panic(err)
	}
	fmt.Println(superheroes)

	return superheroes
}
