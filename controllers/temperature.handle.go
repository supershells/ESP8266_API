package controllers

import (
	"esp8266_api/services"
	valid "esp8266_api/util/validator"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type temperatureHandler struct {
	tempSrv services.TemperatureService
}

func NewTemperatureHandler(tempSrv services.TemperatureService) TemperatureHandler {
	return temperatureHandler{tempSrv: tempSrv}
}

func (h temperatureHandler) GetTemperatures(c *fiber.Ctx) error {
	temperatures, err := h.tempSrv.GetTemperatures()
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err)
	}

	return c.JSON(temperatures)
}

func (h temperatureHandler) GetTemperature(c *fiber.Ctx) error {
	temperatureID := c.Params("temperatureID")

	temperature, err := h.tempSrv.GetTemperature(temperatureID)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err)
	}

	return c.JSON(temperature)
}

func (h temperatureHandler) NewTemperature(c *fiber.Ctx) error {
	request := services.NewTemperatureRequest{}
	if err := valid.ParseBodyAndValidate(c, &request); err != nil {
		return c.Status(http.StatusBadRequest).JSON((err))
	}

	temperature, err := h.tempSrv.NewTemperature(request)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": " temperature created successfully",
		"data":    temperature,
	})
}

func (h temperatureHandler) DeleteTemperature(c *fiber.Ctx) error {
	temperatureID := c.Params("temperatureID")

	err := h.tempSrv.DeleteTemperature(temperatureID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	return c.Status(http.StatusFound).JSON(fiber.Map{
		"message": temperatureID + " was deleted",
	})
}
