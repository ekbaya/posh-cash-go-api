package controllers

import (
	"posh-pesa-api/database"
	"posh-pesa-api/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const SecreteKey = "secret"

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
		Phone:    data["phone"],
		Country:  data["country"],
	}

	database.DB.Create(&user)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "0",
		"sucess":  "true",
		"message": "User created",
	})
}

func Profile(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"Message": "We are here",
	})
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"status":  "1",
			"success": "false",
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  "1",
			"success": "false",
			"message": "Incorrect password",
		})
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(SecreteKey))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"sucess":  "true",
		"status":  "0",
		"message": "Logged in sucessifully",
		"user":    user,
		"token":   t,
	})

}
