package report

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func ReportRoutes(r *gin.Engine, db *gorm.DB) {
    repo := NewGormReportRepository(db)
    service := NewReportService(repo)
    handler := NewReportHandler(service)

    reportGroup := r.Group("/api/reports")
    {
        reportGroup.POST("", handler.CreateReport)
        reportGroup.GET("", handler.ListReports)
        reportGroup.GET("/:id", handler.GetReport)
        reportGroup.PUT("/:id", handler.UpdateReport)
        reportGroup.DELETE("/:id", handler.DeleteReport)
    }
}
