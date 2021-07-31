package main

import (
	"posh-pesa-api/database"
	"posh-pesa-api/routes"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()
	routes.Setup(app)
	govalidator.SetFieldsRequiredByDefault(true)
	app.Listen(":8000")
}
