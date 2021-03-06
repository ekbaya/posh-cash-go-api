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
	app.Get("/api/user/income-range", controllers.FetchIncomeRanges)
	app.Get("/api/user/marital-status", controllers.FetchMaitalStatuses)

	//Protected APIs
	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		ErrorHandler: func(c *fiber.Ctx, e error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "0",
				"sucess":  "false",
				"message": "Unauthorised Access",
			})
		},
		SigningKey: []byte("secret"),
	}))
	app.Get("/api/user/profile", controllers.Profile)
	app.Post("/api/user/profile/update-image", controllers.UpdateUserProfileImage)
	app.Post("/api/user/income-range", controllers.CreateRange)
	app.Post("/api/user/marital-status", controllers.CreateStatus)
}
