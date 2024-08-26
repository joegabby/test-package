package routes

import (
	"github.com/gofiber/fiber/v2"
)

func MainRouter(app fiber.Router) {
	indexRoute := app.Group("api/v1/")

	userroutes(indexRoute)
	Authroutes(indexRoute)
}
