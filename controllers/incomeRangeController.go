package controllers

import (
	"posh-pesa-api/database"
	"posh-pesa-api/models"

	"github.com/gofiber/fiber/v2"
)

func CreateRange(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	income := models.Income{
		Description: data["description"],
	}

	database.DB.Create(&income)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "0",
		"sucess":  "true",
		"message": "Income range created successfully",
	})
}
