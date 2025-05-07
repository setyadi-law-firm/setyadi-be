package report

type CreateReportRequest struct {
    Title   string `json:"title" binding:"required"`
    Content string `json:"content" binding:"required"`
}

type UpdateReportRequest struct {
    Title   string `json:"title"`
    Content string `json:"content"`
}

type ReportResponse struct {
    ID        string `json:"id"`
    Title     string `json:"title"`
    Content   string `json:"content"`
    AuthorID  uint   `json:"author_id"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}
