package database

import (
	"posh-pesa-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:@/go_database"), &gorm.Config{})

	if err != nil {
		panic("Could not establish a connection to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{}, &models.Income{})
}
