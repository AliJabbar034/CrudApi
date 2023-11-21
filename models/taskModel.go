package models

import (
	"context"
	"log"
	"time"

	"githu.com/alijabbar/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var taskColl *mongo.Collection

type Tasks struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty"`
	Description string             `json:"description,omitempty" `
	CreatedAt   time.Time          `json:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at,omitempty"`
	User        User               `json:"user,omitempty"`
}

func init() {
	_, taskColl = config.GetCollection()
}

func (task *Tasks) TaskCreate() (string, error) {

	insertedId, err := taskColl.InsertOne(context.Background(), &task)
	if err != nil {
		return "", err
	}
	id := insertedId.InsertedID.(primitive.ObjectID)

	return id.Hex(), nil

}

func AllTask() ([]Tasks, error) {
	var tasks []Tasks
	cursor, err := taskColl.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}
	if err := cursor.All(context.Background(), &tasks); err != nil {
		return nil, err
	}

	return tasks, nil

}

func (task *Tasks) GetATasks(id string) error {
	taskId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": taskId}

	err := taskColl.FindOne(context.Background(), filter).Decode(&task)
	return err
}

func DeleteOne(id string) (int64, error) {
	taskId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": taskId}

	start := time.Now()
	res, err := taskColl.DeleteOne(context.Background(), filter)
	elapsed := time.Since(start)
	log.Printf("DeleteOne took %s", elapsed)

	return res.DeletedCount, err
}

func UpdateOne(id string, task Tasks) (int64, error) {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": task}

	updated, err := taskColl.UpdateOne(context.Background(), filter, update)

	return updated.ModifiedCount, err

}
