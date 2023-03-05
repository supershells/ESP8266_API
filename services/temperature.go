package services

import "time"

type NewTemperatureRequest struct {
	TempID         string    `json:"-"`
	TempCelsius    float64   `json:"temp_celsius" validate:"required"`
	TempFahrenheit float64   `json:"temp_fahrenheit"`
	Machine        string    `json:"machine" validate:"required"`
	Date           string    `json:"date"`
	Status         int       `json:"status"`
	CreateAt       time.Time `json:"create_at"`
	UpdateAt       time.Time `json:"update_at"`
}

type TemperatureResponse struct {
	TempID         string    `json:"temp_id"`
	TempCelsius    float64   `json:"temp_celsius"`
	TempFahrenheit float64   `json:"temp_fahrenheit"`
	Machine        string    `json:"machine"`
	Date           string    `json:"date"`
	CreateAt       time.Time `json:"create_at"`
	UpdateAt       time.Time `json:"update_at"`
}

type TemperatureService interface {
	GetTemperatures() ([]TemperatureResponse, error)
	GetTemperature(string) (*TemperatureResponse, error)
	NewTemperature(NewTemperatureRequest) (*TemperatureResponse, error)
	DeleteTemperature(string) error
}
