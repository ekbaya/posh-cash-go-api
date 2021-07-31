package controllers

import (
	"posh-pesa-api/database"
	"posh-pesa-api/models"

	"github.com/gofiber/fiber/v2"
)

func CreateStatus(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	maritalStatus := models.MaritalStatus{
		Description: data["description"],
	}

	database.DB.Create(&maritalStatus)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "0",
		"success": "true",
		"message": "Marital status created successfully",
	})
}

func FetchMaitalStatuses(c *fiber.Ctx) error {
	var marital_statuses []models.MaritalStatus

	database.DB.Find(&marital_statuses)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "0",
		"sucess":  "true",
		"message": "Marital statuses fetched sucessfully",
		"data":    marital_statuses,
	})
}
