package services

import (
	"context"
	"encoding/json"
	"esp8266_api/errs"
	"esp8266_api/logs"
	"esp8266_api/repositories"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type temperatureSeriveRedis struct {
	tempRepo    repositories.TemperatureRepository
	redisClient *redis.Client
}

func NewTemperatureServiceRedis(tempRepo repositories.TemperatureRepository, redisClient *redis.Client) TemperatureService {
	return temperatureSeriveRedis{tempRepo, redisClient}
}

func (s temperatureSeriveRedis) GetTemperatures() (tempResponses []TemperatureResponse, err error) {

	key := "service::GetTemperatures"
	if TemperatureJson, err := s.redisClient.Get(context.Background(), key).Result(); err == nil {
		if json.Unmarshal([]byte(TemperatureJson), &tempResponses) == nil {
			fmt.Println("redis")
			return tempResponses, nil
		}
	}

	temperature, err := s.tempRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	tempResponses = []TemperatureResponse{}
	for _, temperature := range temperature {
		tempResponse := TemperatureResponse{
			TempID:         temperature.TempID,
			TempCelsius:    temperature.TempCelsius,
			TempFahrenheit: temperature.TempFahrenheit,
			Machine:        temperature.Machine,
			Date:           temperature.Date,
			CreateAt:       temperature.CreateAt,
			UpdateAt:       temperature.UpdateAt,
		}
		tempResponses = append(tempResponses, tempResponse)
	}

	//Redis SET
	if data, err := json.Marshal(tempResponses); err == nil {
		s.redisClient.Set(context.Background(), key, string(data), time.Second*10)
	}

	fmt.Println("database")

	return tempResponses, err
}

func (s temperatureSeriveRedis) GetTemperature(id string) (*TemperatureResponse, error) {
	temperature, err := s.tempRepo.GetById(id)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	tempResponse := TemperatureResponse{
		TempID:         temperature.TempID,
		TempCelsius:    temperature.TempCelsius,
		TempFahrenheit: temperature.TempFahrenheit,
		Machine:        temperature.Machine,
		Date:           temperature.Date,
		CreateAt:       temperature.CreateAt,
		UpdateAt:       temperature.UpdateAt,
	}

	return &tempResponse, nil
}

func (s temperatureSeriveRedis) NewTemperature(request NewTemperatureRequest) (*TemperatureResponse, error) {

	location, _ := time.LoadLocation("Asia/Bangkok")
	fahrenheit := (1.8 * (request.TempCelsius)) + 32
	today := time.Now().Format("2006-01-02")
	temperature := repositories.Temperature{
		TempID:         uuid.New().String(),
		TempCelsius:    request.TempCelsius,
		TempFahrenheit: fahrenheit,
		Machine:        request.Machine,
		Status:         1,
		Date:           today,
		CreateAt:       time.Now().In(location),
		UpdateAt:       time.Now().In(location),
	}

	newTemperature, err := s.tempRepo.Create(temperature)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	response := TemperatureResponse{
		TempID:         newTemperature.TempID,
		TempCelsius:    newTemperature.TempCelsius,
		TempFahrenheit: newTemperature.TempFahrenheit,
		Machine:        newTemperature.Machine,
		Date:           newTemperature.Date,
		CreateAt:       newTemperature.CreateAt,
		UpdateAt:       newTemperature.UpdateAt,
	}

	return &response, nil
}

func (s temperatureSeriveRedis) DeleteTemperature(id string) error {
	err := s.tempRepo.Delete(id)
	if err != nil {
		logs.Error(err)
		return errs.NewUnexpectedError()
	}
	return nil
}
