package config

import (
	"github.com/ternaknesia/qolbu/exception"
	"github.com/gofiber/fiber/v2"
)

func FiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
