package controllers

import(
	"github.com/gofiber/fiber/v2"
)

func Testt(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message":"upojoi",
	})
}