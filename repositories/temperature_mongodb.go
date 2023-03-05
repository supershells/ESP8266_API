package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type temperatureRepositoryMongoDB struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewtemperatureRepositoryMongoDB(client *mongo.Client, db *mongo.Database) TemperatureRepository {
	return &temperatureRepositoryMongoDB{client, db}
}

func (r *temperatureRepositoryMongoDB) GetAll() ([]Temperature, error) {
	var temperatures []Temperature
	tem, err := r.db.Collection("temperature").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer tem.Close(context.Background())
	for tem.Next(context.Background()) {
		var temperature Temperature
		if err := tem.Decode(&temperature); err != nil {
			return nil, err
		}
		temperatures = append(temperatures, temperature)
	}
	if err := tem.Err(); err != nil {
		return nil, err
	}

	return temperatures, nil
}

func (r *temperatureRepositoryMongoDB) GetById(id string) (*Temperature, error) {
	filter := bson.M{"temp_id": id}
	var temperature Temperature
	err := r.db.Collection("temperature").FindOne(context.Background(), filter).Decode(&temperature)
	if err != nil {
		return nil, err
	}
	return &temperature, nil
}

func (r *temperatureRepositoryMongoDB) GetByMachine(machine string) ([]Temperature, error) {
	filter := bson.M{"machine": machine}
	var temperatures []Temperature
	tem, err := r.db.Collection("temperature").Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer tem.Close(context.Background())
	for tem.Next(context.Background()) {
		var temperature Temperature
		if err := tem.Decode(&temperature); err != nil {
			return nil, err
		}
		temperatures = append(temperatures, temperature)
	}
	if err := tem.Err(); err != nil {
		return nil, err
	}
	return temperatures, nil
}

func (r *temperatureRepositoryMongoDB) Create(temperature Temperature) (*Temperature, error) {
	_, err := r.db.Collection("temperature").InsertOne(context.Background(), temperature)
	if err != nil {
		return nil, err
	}
	return &temperature, nil
}

func (r *temperatureRepositoryMongoDB) Delete(id string) error {
	filter := bson.M{"temp_id": id}
	_, err := r.db.Collection("temperature").DeleteOne(context.Background(), filter)
	return err
}
