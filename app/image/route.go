package image

import (
    "github.com/gin-gonic/gin"
    "github.com/setyadi-law-firm/setyadi-be/app/models"
)

func ImageRoutes(r *gin.Engine, config *models.Config) {
    supabase := NewSupabase(config)
    service := NewImageService(supabase)
    handler := NewImageHandler(service)

    imageGroup := r.Group("/images")
    {
        imageGroup.POST("", handler.UploadImage)
    }
}
