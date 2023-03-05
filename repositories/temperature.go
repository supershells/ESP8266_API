package repositories

import "time"

type Temperature struct {
	TempID         string    `bson:"temp_id"`
	TempCelsius    float64   `bson:"temp_celsius"`
	TempFahrenheit float64   `bson:"temp_fahrenheit"`
	Machine        string    `bson:"machine"`
	Status         int       `bson:"status"` // 0 = normal, 1 = warning, 2 = danger
	Date           string    `bson:"date"`
	CreateAt       time.Time `bson:"create_at"`
	UpdateAt       time.Time `bson:"update_at"`
}

type TemperatureRepository interface {
	GetAll() ([]Temperature, error)
	GetById(string) (*Temperature, error)
	GetByMachine(string) ([]Temperature, error)
	Create(Temperature) (*Temperature, error)
	//Update(*Temperature) (*Temperature, error)
	Delete(string) error
}
