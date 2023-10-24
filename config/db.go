package config

import (
	"api-gin/models"
	"api-gin/utils"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	username := utils.Getenv("MYSQLUSER", "root")
	password := utils.Getenv("MYSQLPASSWORD", "password")
	host := utils.Getenv("MYSQLHOST", "127.0.0.1")
	port := utils.Getenv("MYSQLPORT", "3306")
	database := utils.Getenv("MYSQLDATABASE", "database_movie")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.Movie{}, &models.AgeRatingCategory{}, &models.User{})

	return db
}
