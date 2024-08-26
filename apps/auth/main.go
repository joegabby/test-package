package main

import (
	// "bufio"
	// "fmt"
	// "os"
	// "io/fs"
	// "sort"

	"github.com/gofiber/fiber/v2"
	"github.com/joegabby/ogla-backend.git/helpers"

	"github.com/joegabby/ogla-backend.git/apps/auth/routes"
	_ "github.com/joho/godotenv"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/gorm"
)

func main() {
	// helpers.Verify
	// helpers.DbInit()
	// shared.Verify()
	app := fiber.New()
	mainRoute := app.Group("/")
	routes.MainRouter(mainRoute)
	app.Listen(":9000")
}
