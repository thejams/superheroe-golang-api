//Package repository provides all the methods to persit data in the DB
package repository

import (
	"context"
	"fmt"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"superheroe-api/superheroe-golang-api/src/config"
	"superheroe-api/superheroe-golang-api/src/entity"
	"superheroe-api/superheroe-golang-api/src/util"
)

var (
	once sync.Once
)

// mongoDB main struct for mongodb logic
type mongoDB struct {
	db *mongo.Database
}

// NewMongoDB returns a mongodb struct
func NewMongoDB() *mongoDB {
	return &mongoDB{}
}

// Conn provides a new mongodb connection
func (m *mongoDB) Conn(ctx context.Context, cfg *config.APPConfig) error {
	var err error
	log.SetFormatter(&log.JSONFormatter{})

	once.Do(func() {
		uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", cfg.MONGO_USER, cfg.MONGO_PWD, cfg.MONGO_HOST, cfg.MONGO_PORT, cfg.MONGO_DB)

		client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
		if err != nil {
			log.WithFields(log.Fields{"package": "repository", "client": "MongoDB", "method": "NewMongoConnection"}).Error(err.Error())
		}

		database := client.Database(cfg.MONGO_DB)

		cancel_ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		err = client.Ping(cancel_ctx, readpref.Primary())
		if err != nil {
			log.WithFields(log.Fields{"package": "repository", "client": "MongoDB", "method": "NewMongoConnection"}).Error(err.Error())
			err = fmt.Errorf("Ping to database failed")
		}

		m.db = database
	})

	if err != nil {
		return err
	}

	log.WithFields(log.Fields{"package": "repository", "client": "MongoDB", "method": "NewMongoConnection"}).Info("MongoDB connected")
	return nil
}

func (m *mongoDB) Close(ctx context.Context) error {
	/* if err := client.Disconnect(ctx); err != nil {
		panic(err)
	} */
	return m.db.Client().Disconnect(ctx)
}

//GetAll returns all the superheroes in the DB
func (r *mongoDB) GetAll(ctx context.Context) ([]entity.Character, error) {
	var superheroes []entity.Character
	collection := r.db.Collection("superheroe")
	filter := bson.M{}

	cancel_ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(cancel_ctx, filter)
	if err != nil {
		log.WithFields(log.Fields{"package": "repository", "client": "MongoDB", "method": "GetAll"}).Error(err.Error())
		return nil, err
	}
	defer cursor.Close(cancel_ctx)

	if err = cursor.All(cancel_ctx, &superheroes); err != nil {
		log.WithFields(log.Fields{"package": "repository", "client": "MongoDB", "method": "GetAll"}).Error(err.Error())
		return nil, err
	}

	log.WithFields(log.Fields{"package": "repository", "client": "MongoDB", "method": "GetAll"}).Info("ok")
	return superheroes, nil
}

//Get returns a single superheroe from the DB
func (r *mongoDB) Get(i string, ctx context.Context) (*entity.Character, error) {
	var result *entity.Character
	collection := r.db.Collection("superheroe")
	oid, err := primitive.ObjectIDFromHex(i)
	if err != nil {
		log.WithFields(log.Fields{"package": "repository", "client": "MongoDB", "method": "Get"}).Error(err.Error())
		return nil, err
	}

	cancel_ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": oid}
	err = collection.FindOne(cancel_ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		log.WithFields(log.Fields{"package": "repository", "client": "MongoDB", "method": "Get"}).Error(err.Error())
		return nil, &util.NotFoundError{Message: fmt.Sprintf("no superheroe with id %v found", oid)}
	} else if err != nil {
		log.WithFields(log.Fields{"package": "repository", "client": "MongoDB", "method": "Get"}).Error(err.Error())
		return nil, err
	}

	log.WithFields(log.Fields{"package": "repository", "client": "MongoDB", "method": "Get"}).Info("ok")
	return result, nil
}

//Add add a new superheroe to the DB
func (r *mongoDB) Add(c *entity.Character, ctx context.Context) (*entity.Character, error) {
	collection := r.db.Collection("superheroe")
	filter := bson.D{{"name", c.Name}, {"alias", c.Alias}}

	cancel_ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := collection.InsertOne(cancel_ctx, filter)
	if err != nil {
		log.WithFields(log.Fields{"package": "repository", "client": "MongoDB", "method": "Add"}).Error(err.Error())
		return nil, err
	}

	id := res.InsertedID.(primitive.ObjectID)
	c.ID = id

	log.WithFields(log.Fields{"package": "repository", "client": "MongoDB", "method": "Add"}).Info("ok")
	return c, nil
}

//Delete remove a superheroe from the DB
func (r *mongoDB) Delete(id string, ctx context.Context) (string, error) {
	collection := r.db.Collection("superheroe")
	oid, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": oid}

	cancel_ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := collection.DeleteOne(cancel_ctx, filter)
	if err != nil {
		log.WithFields(log.Fields{"package": "repository", "client": "MongoDB", "method": "Delete"}).Error(err.Error())
		return "", err
	}
	if res.DeletedCount == 0 {
		err := fmt.Sprintf("document not found %v", res)
		log.WithFields(log.Fields{"package": "repository", "client": "MongoDB", "method": "Delete"}).Error(err)
		return "", &util.NotFoundError{Message: err}
	}

	log.WithFields(log.Fields{"package": "repository", "client": "MongoDB", "method": "Delete"}).Info("ok")
	return fmt.Sprintf("document deleted %v", res), nil
}

//Edit updates a superheroe in DB with new information
func (r *mongoDB) Edit(id string, c *entity.Character, ctx context.Context) (*entity.Character, error) {
	collection := r.db.Collection("superheroe")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.WithFields(log.Fields{"package": "repository", "client": "MongoDB", "method": "Edit"}).Error(err)
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
		log.WithFields(log.Fields{"package": "repository", "client": "MongoDB", "method": "Edit"}).Error(err)
		return nil, err
	}
	c.ID = oid

	log.WithFields(log.Fields{"package": "repository", "client": "MongoDB", "method": "Edit"}).Info("ok")
	return c, nil
}
