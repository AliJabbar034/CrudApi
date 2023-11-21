package models

import (
	"context"
	"fmt"

	"githu.com/alijabbar/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var coll *mongo.Collection

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FName    string             `json:"firstName"`
	LName    string             `json:"lastName"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}

func init() {
	coll, _ = config.GetCollection()

}

func CreateUser(user User) (string, error) {

	inserted, er := coll.InsertOne(context.TODO(), &user)
	if er != nil {
		return "", er
	}

	userId, ok := inserted.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("failed to convert InsertedID to primitive.ObjectID")
	}
	id := userId.Hex()

	return id, nil

}

func FindById(id string) (User, error) {

	var user User
	d, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": d}
	res := coll.FindOne(context.Background(), filter).Decode(&user)

	return user, res

}

func Login(email string) (User, error) {
	var user User
	filter := bson.M{"email": email}

	res := coll.FindOne(context.TODO(), filter).Decode(&user)

	return user, res
}

func UserUPdate(user User) (int64, error) {
	id := user.ID
	filter := bson.M{"_id": id}
	update := bson.M{"$set": user}

	updated, err := coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return 0, err
	}

	return updated.ModifiedCount, nil
}
