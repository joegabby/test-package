package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joegabby/ogla-backend.git/apps/auth/controllers"
)

func userroutes(app fiber.Router) {
	user:=app.Group("/users")

	user.Get("/me", controllers.Testt)
}
