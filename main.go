package main

import (
	"posh-pesa-api/database"
	"posh-pesa-api/routes"

	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/jwt/v2"
)

func main() {
	database.Connect()
	app := fiber.New()
	routes.Setup(app)
	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	app.Listen(":8000")
}
