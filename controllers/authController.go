package controllers

import (
	"posh-pesa-api/database"
	"posh-pesa-api/models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const SecreteKey = "secrete"

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
	return c.JSON(user)
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

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(SecreteKey))

	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"sucess":  "false",
			"status":  "1",
			"message": "Could not login",
		})
	}

	return c.JSON(fiber.Map{
		"sucess":  "true",
		"status":  "0",
		"message": "Logged in sucessifully",
		"user":    user,
		"token":   token,
	})
}
