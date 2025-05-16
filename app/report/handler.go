package report

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

type ReportHandler struct {
    service ReportService
}

func NewReportHandler(service ReportService) *ReportHandler {
    return &ReportHandler{service}
}

func (h *ReportHandler) CreateReport(c *gin.Context) {
    var input CreateReportRequest
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Sementara authorID dummy 1
    authorID := uint(1)

    report, err := h.service.CreateReport(input, authorID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, report)
}

func (h *ReportHandler) GetReport(c *gin.Context) {
    idParam := c.Param("id")
    id, err := uuid.Parse(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
        return
    }

    report, err := h.service.GetReport(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Report not found"})
        return
    }

    c.JSON(http.StatusOK, report)
}

func (h *ReportHandler) UpdateReport(c *gin.Context) {
    idParam := c.Param("id")
    id, err := uuid.Parse(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
        return
    }

    var input UpdateReportRequest
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    report, err := h.service.UpdateReport(id, input)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, report)
}

func (h *ReportHandler) DeleteReport(c *gin.Context) {
    idParam := c.Param("id")
    id, err := uuid.Parse(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
        return
    }

    if err := h.service.DeleteReport(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Report deleted successfully"})
}

func (h *ReportHandler) BulkDeleteReports(c *gin.Context) {
	var req BulkDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := h.service.BulkDeleteReports(req.ReportIDs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reports deleted successfully"})
}

func (h *ReportHandler) ListReports(c *gin.Context) {
    reports, err := h.service.ListReports()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, reports)
}
