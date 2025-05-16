package report

import (
    "github.com/setyadi-law-firm/setyadi-be/app/auth"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func ReportRoutes(r *gin.Engine, db *gorm.DB, authUtil *auth.Util) {
    repo := NewGormReportRepository(db)
    service := NewReportService(repo)
    handler := NewReportHandler(service)

    reportGroup := r.Group("/api/reports")
    reportGroup.Use(authUtil.JwtAuthMiddleware())
    {
        reportGroup.POST("", handler.CreateReport)
        reportGroup.GET("", handler.ListReports)
        reportGroup.GET("/:id", handler.GetReport)
        reportGroup.PUT("/:id", handler.UpdateReport)
        reportGroup.DELETE("/bulk", handler.BulkDeleteReports)
        reportGroup.DELETE("/:id", handler.DeleteReport)
    }
}
