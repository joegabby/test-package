package routes

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/joegabby/ogla-backend.git/controllers"
	// "github.com/joegabby/ogla-backend.git/"
)

func Authroutes(app fiber.Router) {

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.JSON(
			"m",
		)
	})
}
