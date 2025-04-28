package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"

	"github.com/setyadi-law-firm/setyadi-be/app/auth"
	"github.com/setyadi-law-firm/setyadi-be/app/image"
	"github.com/setyadi-law-firm/setyadi-be/app/models"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found. Using default environment values.")
	}

	config := models.LoadConfig()

	db, err := gorm.Open(postgres.Open(config.Dsn()), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to PostgreSQL:", err)
	}

	if err := db.AutoMigrate(&auth.User{}); err != nil {
		log.Fatal("failed to migrate user model:", err)
	}

	r := gin.Default()

	auth.AuthRoutes(r, db, config)
	image.ImageRoutes(r, config)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("failed to start server:", err)
	}
}
