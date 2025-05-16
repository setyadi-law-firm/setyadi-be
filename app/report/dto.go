package report

import "github.com/google/uuid"

type CreateReportRequest struct {
    Title   string `json:"title" binding:"required"`
    Content string `json:"content" binding:"required"`
    Author string `json:"author" binding:"required"`
    ImageURL string `json:"image_url" binding:"required"`
}

type UpdateReportRequest struct {
    Title   string `json:"title"`
    Content string `json:"content"`
    ImageURL string `json:"image_url"`
}

type BulkDeleteRequest struct {
	ReportIDs []uuid.UUID `json:"report_ids"`
}

type ReportResponse struct {
    ID        string `json:"id"`
    Title     string `json:"title"`
    Content   string `json:"content"`
    Author  string   `json:"author"`
    ImageURL string `json:"image_url" binding:"required"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}
