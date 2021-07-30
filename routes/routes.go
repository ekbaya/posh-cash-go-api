package routes

import (
	"posh-pesa-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/user/register", controllers.Register)
	app.Post("/api/user/login", controllers.Login)
	app.Get("/api/user/profile", controllers.Profile)
}
