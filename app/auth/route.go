package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/setyadi-law-firm/setyadi-be/app/models"
)

func AuthRoutes(r *gin.Engine, db *gorm.DB, config *models.Config) {
	util := NewUtil(config)
	repo := NewGormAuthRepository(db)
	service := NewAuthService(repo, util)
	handler := NewAuthHandler(service)

	authGroup := r.Group("/api/auth")
	{
		authGroup.POST("/register", handler.Register)
		authGroup.POST("/login", handler.Login)
	}
}
