package services

import (
	"esp8266_api/errs"
	"esp8266_api/logs"
	"esp8266_api/repositories"
	"time"

	"github.com/google/uuid"
)

type temperatureSerive struct {
	tempRepo repositories.TemperatureRepository
}

func NewTemperatureService(tempRepo repositories.TemperatureRepository) TemperatureService {
	return temperatureSerive{tempRepo: tempRepo}
}

func (s temperatureSerive) GetTemperatures() ([]TemperatureResponse, error) {
	temperature, err := s.tempRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewNotFoundError("Temperature not found")
	}

	tempResponses := []TemperatureResponse{}
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
	return tempResponses, nil
}

func (s temperatureSerive) GetTemperature(id string) (*TemperatureResponse, error) {
	temperature, err := s.tempRepo.GetById(id)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewNotFoundError("Temperature not found")
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

func (s temperatureSerive) NewTemperature(request NewTemperatureRequest) (*TemperatureResponse, error) {

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

func (s temperatureSerive) DeleteTemperature(id string) error {
	err := s.tempRepo.Delete(id)
	if err != nil {
		logs.Error(err)
		return errs.NewUnexpectedError()
	}
	return nil
}
