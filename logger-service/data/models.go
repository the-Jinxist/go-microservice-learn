package data

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func New(mongo *mongo.Client) Models {
	client = mongo
	return Models{
		LogEntry: LogEntry{},
	}
}

type Models struct {
	LogEntry LogEntry
}

type LogEntry struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string    `bson:"name" json:"name"`
	Data      string    `bson:"data" json:"data"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

func (l *LogEntry) Insert(entry LogEntry) error {
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), LogEntry{
		Name:      entry.Name,
		Data:      entry.Data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		log.Printf("error inserting into log: %s", err)
		return err
	}

	return err
}

func (l *LogEntry) All() ([]*LogEntry, error) {
	cxt, canc := context.WithTimeout(context.Background(), time.Second*15)
	defer canc()

	collection := client.Database("logs").Collection("logs")

	opts := options.Find()
	opts.SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := collection.Find(context.TODO(), opts)
	if err != nil {
		log.Printf("error finding all logs: %s", err)
		return nil, err
	}

	defer cursor.Close(cxt)

	var logs []*LogEntry
	for cursor.Next(cxt) {
		var item LogEntry
		err = cursor.Decode(&item)

		if err != nil {
			log.Printf("error decoding log into slice: %s", err)
			return nil, err
		}

		logs = append(logs, &item)
	}

	return logs, nil

}

func (l *LogEntry) GetOne(id string) (*LogEntry, error) {
	cxt, canc := context.WithTimeout(context.Background(), time.Second*15)
	defer canc()

	collection := client.Database("logs").Collection("logs")

	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("error converting id to object id: %s", err)
		return nil, err
	}

	var entry LogEntry
	err = collection.FindOne(cxt, bson.M{"_id": docID}).Decode(entry)
	if err != nil {
		log.Printf("error finding and decoding value: %s", err)
		return nil, err
	}

	return &entry, nil

}

func (l *LogEntry) DropCollection(collectionNames ...string) error {
	finalCollectionName := "logs"
	if len(collectionNames) > 0 {
		finalCollectionName = collectionNames[0]
	}

	cxt, canc := context.WithTimeout(context.Background(), time.Second*15)
	defer canc()

	collection := client.Database("logs").Collection(finalCollectionName)

	if err := collection.Drop(cxt); err != nil {
		return err
	}

	return nil
}

func (l *LogEntry) Update() (*mongo.UpdateResult, error) {
	cxt, canc := context.WithTimeout(context.Background(), time.Second*15)
	defer canc()

	collection := client.Database("logs").Collection("logs")
	docID, err := primitive.ObjectIDFromHex(l.ID)
	if err != nil {
		log.Printf("error converting id to object id: %s", err)
		return nil, err
	}

	result, err := collection.UpdateByID(
		cxt,
		bson.M{"_id": docID},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "name", Value: l.Name},
				{Key: "data", Value: l.Data},
				{Key: "updated_at", Value: time.Now()},
			},
			},
		},
	)

	if err != nil {
		log.Printf("error updating collection id %s", err)
		return nil, err
	}

	return result, nil

}
