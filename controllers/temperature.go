package controllers

import "github.com/gofiber/fiber/v2"

type TemperatureHandler interface {
	GetTemperatures(c *fiber.Ctx) error
	GetTemperature(c *fiber.Ctx) error
	NewTemperature(c *fiber.Ctx) error
	DeleteTemperature(c *fiber.Ctx) error
}
