package routes

import (
	"posh-pesa-api/controllers"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func Setup(app *fiber.App) {
	//public APIs
	app.Post("/api/user/register", controllers.Register)
	app.Post("/api/user/login", controllers.Login)

	//Protexted APIs
	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	app.Get("/api/user/profile", controllers.Profile)
}
