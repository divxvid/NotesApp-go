package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserPass struct {
	Username string `bson:"username,omitempty"`
	Password string `bson:"password,omitempty"`
}

type Note struct {
	Username string `bson:"username,omitempty"`
	Title    string `bson:"title,omitempty"`
	Body     string `bson:"note,omitempty"`
}

type DBService interface {
	FetchUserByUsername(string) (UserPass, error)
	CreateUser(*UserPass) (string, error)
	CreateNote(*Note) (string, error)
	FetchAllNotes() ([]Note, error)
	FetchNoteById(string) (*Note, error)
}

type MongoDBService struct {
	ctx    *context.Context
	client *mongo.Client
	db     *mongo.Database
}

func (mongoSvc *MongoDBService) FetchNoteById(id string) (*Note, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	collection := mongoSvc.db.Collection("notes")
	var note Note

	query := bson.D{{"_id", objectId}}
	collection.FindOne(*mongoSvc.ctx, query).Decode(&note)

	return &note, nil
}

func (mongoSvc *MongoDBService) CreateNote(note *Note) (string, error) {
	collection := mongoSvc.db.Collection("notes")

	result, err := collection.InsertOne(*mongoSvc.ctx, note)
	if err != nil {
		return "", err
	}
	insertedID := result.InsertedID.(primitive.ObjectID).Hex()
	return insertedID, nil
}

func (mongoSvc *MongoDBService) FetchUserByUsername(username string) (UserPass, error) {
	collection := mongoSvc.db.Collection("userpasses")
	var up UserPass

	query := bson.D{{"username", username}}
	collection.FindOne(*mongoSvc.ctx, query).Decode(&up)
	return up, nil
}

func (mongoSvc *MongoDBService) CreateUser(user *UserPass) (string, error) {
	collection := mongoSvc.db.Collection("userpasses")
	result, err := collection.InsertOne(*mongoSvc.ctx, user)
	if err != nil {
		return "", err
	}
	insertedID := result.InsertedID.(primitive.ObjectID).Hex()
	return insertedID, nil
}

func (mongoSvc *MongoDBService) Close() error {
	err := mongoSvc.client.Disconnect(*mongoSvc.ctx)
	return err
}

func NewMongoDBService() (*MongoDBService, error) {
	mongoUsername := os.Getenv("MONGO_USER")
	mongoPassword := os.Getenv("MONGO_PASSWORD")
	mongoCluster := os.Getenv("MONGO_CLUSTER")

	mongoURI := fmt.Sprintf(
		"mongodb+srv://%s:%s@%s",
		mongoUsername, mongoPassword, mongoCluster)

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	db := client.Database("test")
	return &MongoDBService{
		ctx:    &ctx,
		client: client,
		db:     db,
	}, nil
}

func (mongoSvc *MongoDBService) GetCollectionNames() ([]string, error) {
	collectionNames, err := mongoSvc.db.ListCollectionNames(*mongoSvc.ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	return collectionNames, nil
}

func (mongoSvc *MongoDBService) GetAllUserPasses() ([]UserPass, error) {
	collection := mongoSvc.db.Collection("userpasses")

	var up []UserPass

	cursor, err := collection.Find(*mongoSvc.ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(*mongoSvc.ctx)

	if err = cursor.All(*mongoSvc.ctx, &up); err != nil {
		return nil, err
	}

	return up, nil
}

func (mongoSvc *MongoDBService) FetchAllNotes() ([]Note, error) {
	collection := mongoSvc.db.Collection("notes")

	var notes []Note

	cursor, err := collection.Find(*mongoSvc.ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(*mongoSvc.ctx)

	if err = cursor.All(*mongoSvc.ctx, &notes); err != nil {
		return nil, err
	}

	return notes, nil
}
