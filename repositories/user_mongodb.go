package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepositoryMongoDB struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewUserRepositoryMongoDB(client *mongo.Client, db *mongo.Database) UserRepository {
	return &userRepositoryMongoDB{client, db}
}

func (r *userRepositoryMongoDB) GetAll() ([]User, error) {
	var users []User
	cur, err := r.db.Collection("users").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close((context.Background()))
	for cur.Next(context.Background()) {
		var user User
		if err := cur.Decode(&user); err != nil {
			return nil, err
			//log.Fatal(err)
		}
		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (r userRepositoryMongoDB) GetById(id string) (*User, error) {
	filter := bson.M{"user_id": id}
	var user User
	err := r.db.Collection("users").FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r userRepositoryMongoDB) GetByUsername(username string) (*User, error) {
	filter := bson.M{"username": username}
	var user User
	err := r.db.Collection("users").FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r userRepositoryMongoDB) GetByDept(dept string) (*User, error) {
	filter := bson.M{"dept": dept}
	var user User
	err := r.db.Collection("users").FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r userRepositoryMongoDB) Create(user User) (*User, error) {
	_, err := r.db.Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r userRepositoryMongoDB) Update(user *User) (*User, error) {
	filter := bson.M{"user_id": user.UserID}
	update := bson.M{"$set": user}
	_, err := r.db.Collection("users").UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r userRepositoryMongoDB) Delete(id string) error {
	filter := bson.M{"user_id": id}
	_, err := r.db.Collection("users").DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}
