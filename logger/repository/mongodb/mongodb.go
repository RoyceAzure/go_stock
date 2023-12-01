package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const LOG_DATABASE = "logs"
const LOG_COLLECTION = "logs"

type IMongoDao interface {
	Insert(ctx context.Context, entry LogEntry) error
	GetAll(ctx context.Context) ([]*LogEntry, error)
	InsertString(ctx context.Context, log string) error
}

type MongoDao struct {
	client *mongo.Client
}

type LogEntry struct {
	ID          string    `bson:"_id,omitempty" json:"id,omitempty"`
	ServiceName string    `bson:"service_name" json:"service_name"`
	Level       string    `bson:"level" json:"level"`
	Message     string    `bson:"message" json:"message"`
	Error       string    `bson:"error" json:"error"`
	CreatedAt   time.Time `bson:"time" json:"time"`
}

func NewMongoDao(client *mongo.Client) IMongoDao {
	return &MongoDao{
		client: client,
	}
}

func (dao *MongoDao) Insert(ctx context.Context, entry LogEntry) error {
	collection := dao.client.Database(LOG_DATABASE).Collection(LOG_COLLECTION)

	_, err := collection.InsertOne(ctx, entry)

	return err
}

func (dao *MongoDao) InsertString(ctx context.Context, log string) error {
	collection := dao.client.Database(LOG_DATABASE).Collection(LOG_COLLECTION)

	record := bson.M{
		"Message": log,
	}

	_, err := collection.InsertOne(ctx, record)

	return err
}

func (dao *MongoDao) GetAll(ctx context.Context) ([]*LogEntry, error) {
	collection := dao.client.Database(LOG_DATABASE).Collection(LOG_COLLECTION)
	opts := options.Find()
	opts.SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := collection.Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var logs []*LogEntry

	for cursor.Next(ctx) {
		var item LogEntry

		err := cursor.Decode(&item)
		if err != nil {
			return nil, err
		} else {
			logs = append(logs, &item)
		}
	}

	return logs, nil
}

func ConnectToMongo(ctx context.Context, address string) (*mongo.Client, error) {
	//自己設定clientOptions  連線資訊
	//透過mongo套件使用自己設定的clientOptions連線
	clientOptions := options.Client().ApplyURI(address)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})
	//初始化連線需要context?
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	return c, nil
}
