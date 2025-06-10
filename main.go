package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"

	"github.com/setyadi-law-firm/setyadi-be/app/auth"
	"github.com/setyadi-law-firm/setyadi-be/app/report"
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
	if err := db.AutoMigrate(&report.Report{}); err != nil {
		log.Fatal("failed to migrate user model:", err)
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"https://www.dssetyadipartners.com",
		},
		AllowHeaders: []string{
			"Origin", "Content-Type", "Authorization",
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH",
		},
		AllowCredentials: true,
	}))

	authUtil := auth.NewUtil(config)

	auth.AuthRoutes(r, db, config)
	image.ImageRoutes(r, config, authUtil)
	report.ReportRoutes(r, db, authUtil)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("failed to start server:", err)
	}
}
