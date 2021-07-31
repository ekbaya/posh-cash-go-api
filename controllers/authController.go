package controllers

import (
	"fmt"
	"posh-pesa-api/database"
	"posh-pesa-api/models"
	"strconv"
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

	var u models.User
	database.DB.Where("email = ?", data["email"]).First(&u)

	if u.Id > 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  "1",
			"success": "false",
			"message": "User with this email already exist",
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		FirstName:        data["first_name"],
		LastName:         data["last_name"],
		Email:            data["email"],
		Password:         password,
		Phone:            data["phone"],
		AlternativePhone: data["alternative_phone"],
		Idnumber:         data["id_number"],
		Dob:              data["dob"],
		Gender:           data["gender"],
		MaritalStatus:    data["marital_status"],
		Address:          data["address"],
		Latitude:         data["latitude"],
		Longitude:        data["longitude"],
		JobType:          data["job_type"],
		Referee:          data["referee"],
		Country:          data["country"],
		Comapny:          data["company"],
		Position:         data["company_position"],
		CompanyPhone:     data["company_phone"],
		IncomeRange:      data["income_range"],
	}

	database.DB.Create(&user)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "0",
		"sucess":  "true",
		"message": "User created",
	})
}

func Profile(c *fiber.Ctx) error {
	var user models.User
	database.DB.Where("id = ?", ClientID(c)).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"status":  "1",
			"success": "false",
			"message": "user not found",
		})
	}

	var income models.Income
	database.DB.Where("id = ?", user.IncomeRange).First(&income)

	var maritalStatus models.MaritalStatus
	database.DB.Where("id = ?", user.MaritalStatus).First(&maritalStatus)

	address := models.Address{
		Name:      user.Address,
		Latitude:  user.Latitude,
		Longitude: user.Longitude,
	}

	company := models.Company{
		Name:        user.Comapny,
		Position:    user.Position,
		CompayPhone: user.CompanyPhone,
	}

	return c.JSON(fiber.Map{
		"status":  "0",
		"success": "true",
		"message": "user fetched successfully",
		"data": fiber.Map{
			"user":           user,
			"address":        address,
			"company":        company,
			"income":         income,
			"marital_status": maritalStatus,
		},
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
	claims["id"] = strconv.Itoa(int(user.Id))
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(SecreteKey))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var income models.Income
	database.DB.Where("id = ?", user.IncomeRange).First(&income)

	var maritalStatus models.MaritalStatus
	database.DB.Where("id = ?", user.MaritalStatus).First(&maritalStatus)

	address := models.Address{
		Name:      user.Address,
		Latitude:  user.Latitude,
		Longitude: user.Longitude,
	}

	company := models.Company{
		Name:        user.Comapny,
		Position:    user.Position,
		CompayPhone: user.CompanyPhone,
	}

	return c.JSON(fiber.Map{
		"sucess":  "true",
		"status":  "0",
		"message": "Logged in sucessifully",
		"data": fiber.Map{
			"user":           user,
			"address":        address,
			"company":        company,
			"income":         income,
			"marital_status": maritalStatus,
		},
		"token": t,
	})

}

func UpdateUserProfileImage(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(422).JSON(fiber.Map{"errors": [1]string{"Failed to update profile image"}})
	}
	c.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))

	var user models.User
	database.DB.Where("id = ?", ClientID(c)).First(&user)

	database.DB.Model(&user).Update("image", file.Filename)
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status":  "0",
		"sucess":  "true",
		"message": "User updated successfully",
	})
}

// Takes a
func ClientID(c *fiber.Ctx) string {
	client := c.Locals("user").(*jwt.Token)
	claims := client.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	return id
}
