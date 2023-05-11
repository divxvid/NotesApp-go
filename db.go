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

type UserPasses struct {
	Username string `bson:"username,omitempty"`
	Password string `bson:"password,omitempty"`
}

type Note struct {
	Username string `bson:"username,omitempty"`
	Title    string `bson:"title,omitempty"`
	Body     string `bson:"note,omitempty"`
}

type DBService interface {
	FetchUserByUsername(string) (UserPasses, error)
	CreateUser(*UserPasses) (*string, error)
}

type MongoDBService struct {
	ctx    *context.Context
	client *mongo.Client
	db     *mongo.Database
}

func (mongoSvc *MongoDBService) CreateUser(user *UserPasses) (*string, error) {
	collection := mongoSvc.db.Collection("userpasses")
	result, err := collection.InsertOne(*mongoSvc.ctx, user)
	if err != nil {
		return nil, err
	}
	insertedID := result.InsertedID.(primitive.ObjectID).String()
	return &insertedID, nil
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

	// databases, err := client.ListDatabaseNames(ctx, bson.M{})
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println(databases)

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

func (mongoSvc *MongoDBService) GetAllUserPasses() ([]UserPasses, error) {
	collection := mongoSvc.db.Collection("userpasses")

	var up []UserPasses

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
