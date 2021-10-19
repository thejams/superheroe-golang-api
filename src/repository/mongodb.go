//Package repository provides all the methods to persit data in the DB
package repository

import (
	"context"
	"fmt"
	"os"
	"strings"
	"superheroe-api/superheroe-golang-api/src/entity"
	"superheroe-api/superheroe-golang-api/src/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongoRepository main struct for mongodb logic
type mongoRepository struct {
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
		host = "127.0.0.1"
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

	database := client.Database("user-test")
	return &mongoRepository{
		db: database,
	}, client
}

//Close close the connection to the MongoDB database
func (r *mongoRepository) Close(ctx context.Context) {
	r.db.Client().Disconnect(ctx)
}

//GetSuperheroes returns all the superheroes in the DB
func (r *mongoRepository) GetSuperheroes(ctx context.Context) ([]*entity.Superhero, error) {
	var superheroes []*entity.Superhero
	collection := r.db.Collection("superheroe")
	filter := bson.M{}
	cursor, err := collection.Find(ctx, filter)
	defer cursor.Close(ctx)

	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &superheroes); err != nil {
		return nil, err
	}
	fmt.Println(superheroes)

	return superheroes, nil
}

//GetSuperheroeById returns a single superheroe from the DB
func (r *mongoRepository) GetSuperheroeById(i string, ctx context.Context) (*entity.Superhero, error) {
	var result *entity.Superhero
	collection := r.db.Collection("superheroe")
	oid, err := primitive.ObjectIDFromHex(i)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": oid}
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, &util.NotFoundError{Message: fmt.Sprintf("no superheroe with id %v found", oid)}
	} else if err != nil {
		return nil, err
	}

	return result, nil
}

//AddSuperheroe add a new superheroe to the DB
func (r *mongoRepository) AddSuperheroe(c *entity.Superhero, ctx context.Context) (*entity.Superhero, error) {
	collection := r.db.Collection("superheroe")
	filter := bson.D{{"name", c.Name}, {"alias", c.Alias}}
	res, err := collection.InsertOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	id := res.InsertedID.(primitive.ObjectID)
	c.ID = id

	return c, nil
}

//DeleteSuperheroe remove a superheroe from the DB
func (r *mongoRepository) DeleteSuperheroe(id string, ctx context.Context) (string, error) {
	collection := r.db.Collection("superheroe")
	oid, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": oid}
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return "", err
	}
	if res.DeletedCount == 0 {
		return "", &util.NotFoundError{Message: fmt.Sprintf("document not found %v", res)}
	}

	return fmt.Sprintf("document deleted %v", res), nil
}

//EditSuperheroe updates a superheroe in DB with new information
func (r *mongoRepository) EditSuperheroe(id string, c *entity.Superhero, ctx context.Context) (*entity.Superhero, error) {
	collection := r.db.Collection("superheroe")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": oid}
	set := bson.D{
		{"$set", bson.D{
			{"name", c.Name},
			{"alias", c.Alias},
		}},
	}
	_, err = collection.UpdateOne(ctx, filter, set)
	if err != nil {
		return nil, err
	}
	c.ID = oid

	return c, nil
}
