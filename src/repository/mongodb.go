//Package repository provides all the methods to persit data in the DB
package repository

import (
	"context"
	"fmt"
	"os"
	"strings"
	"superheroe-api/superheroe-golang-api/src/entity"
	"superheroe-api/superheroe-golang-api/src/util"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

// mongoRepository main struct for mongodb logic
type mongoRepository struct {
	db *mongo.Database
}

// getDBConnection provides all the string connections for the database
func getDBConnection() (string, string, string, string, string) {
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
	db := os.Getenv("MONGO_DB")
	if len(strings.TrimSpace(db)) == 0 {
		db = "superheroes"
	}
	return usr, pwd, host, port, db
}

// NewMongoConnection provides a new mongodb connection
func NewMongoConnection(ctx context.Context) (Repository, error) {
	log.SetFormatter(&log.JSONFormatter{})
	usr, pwd, host, port, db := getDBConnection()
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", usr, pwd, host, port, db)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.WithFields(log.Fields{"package": "repository", "client": "MongoDB", "method": "NewMongoConnection"}).Error(err.Error())
		return nil, err
	}
	database := client.Database(db)

	cancel_ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = client.Ping(cancel_ctx, readpref.Primary())
	if err != nil {
		log.WithFields(log.Fields{"package": "repository", "client": "MongoDB", "method": "NewMongoConnection"}).Error(err.Error())
		return nil, fmt.Errorf("Ping to database failed")
	}

	return &mongoRepository{db: database}, nil
}

func DisconnectDB(ctx context.Context) {
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}

//GetSuperheroes returns all the superheroes in the DB
func (r *mongoRepository) GetSuperheroes(ctx context.Context) ([]entity.Superhero, error) {
	var superheroes []entity.Superhero
	collection := r.db.Collection("superheroe")
	filter := bson.M{}

	cancel_ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(cancel_ctx, filter)
	if err != nil {
		log.WithFields(log.Fields{"package": "repository", "client": "MongoDB", "method": "GetSuperheroes"}).Error(err.Error())
		return nil, err
	}
	defer cursor.Close(cancel_ctx)

	if err = cursor.All(cancel_ctx, &superheroes); err != nil {
		log.WithFields(log.Fields{"package": "repository", "client": "MongoDB", "method": "GetSuperheroes"}).Error(err.Error())
		return nil, err
	}

	log.WithFields(log.Fields{"package": "repository", "client": "MongoDB", "method": "GetSuperheroes"}).Info("ok")
	return superheroes, nil
}

//GetSuperheroeById returns a single superheroe from the DB
func (r *mongoRepository) GetSuperheroeById(i string, ctx context.Context) (*entity.Superhero, error) {
	var result *entity.Superhero
	collection := r.db.Collection("superheroe")
	oid, err := primitive.ObjectIDFromHex(i)
	if err != nil {
		log.WithFields(log.Fields{"package": "repository", "client": "MongoDB", "method": "GetSuperheroeById"}).Error(err.Error())
		return nil, err
	}

	cancel_ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": oid}
	err = collection.FindOne(cancel_ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		log.WithFields(log.Fields{"package": "repository", "client": "MongoDB", "method": "GetSuperheroeById"}).Error(err.Error())
		return nil, &util.NotFoundError{Message: fmt.Sprintf("no superheroe with id %v found", oid)}
	} else if err != nil {
		log.WithFields(log.Fields{"package": "repository", "client": "MongoDB", "method": "GetSuperheroeById"}).Error(err.Error())
		return nil, err
	}

	log.WithFields(log.Fields{"package": "repository", "client": "MongoDB", "method": "GetSuperheroeById"}).Info("ok")
	return result, nil
}

//AddSuperheroe add a new superheroe to the DB
func (r *mongoRepository) AddSuperheroe(c *entity.Superhero, ctx context.Context) (*entity.Superhero, error) {
	collection := r.db.Collection("superheroe")
	filter := bson.D{{"name", c.Name}, {"alias", c.Alias}}

	cancel_ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := collection.InsertOne(cancel_ctx, filter)
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

	cancel_ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := collection.DeleteOne(cancel_ctx, filter)
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

	cancel_ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err = collection.UpdateOne(cancel_ctx, filter, set)
	if err != nil {
		return nil, err
	}
	c.ID = oid

	return c, nil
}
