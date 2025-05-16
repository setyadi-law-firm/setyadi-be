package image

import (
	"github.com/gin-gonic/gin"
	"github.com/setyadi-law-firm/setyadi-be/app/auth"
	"github.com/setyadi-law-firm/setyadi-be/app/models"
)

func ImageRoutes(r *gin.Engine, config *models.Config, authUtil *auth.Util) {
    supabase := NewSupabase(config)
    service := NewImageService(supabase)
    handler := NewImageHandler(service)

    imageGroup := r.Group("/images")
    imageGroup.Use(authUtil.JwtAuthMiddleware())
    {
        imageGroup.POST("", handler.UploadImage)
    }
}
