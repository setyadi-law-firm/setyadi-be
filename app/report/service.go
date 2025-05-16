package report

import (
    "github.com/google/uuid"
)

type ReportService interface {
    CreateReport(input CreateReportRequest, authorID uint) (*Report, error)
    GetReport(id uuid.UUID) (*Report, error)
    UpdateReport(id uuid.UUID, input UpdateReportRequest) (*Report, error)
    DeleteReport(id uuid.UUID) error
    DeleteAllReports() error
    ListReports() ([]*Report, error)
}

type reportService struct {
    repo ReportRepository
}

func NewReportService(repo ReportRepository) ReportService {
    return &reportService{repo}
}

func (s *reportService) CreateReport(input CreateReportRequest, authorID uint) (*Report, error) {
    report := &Report{
        Title:    input.Title,
        Content:  input.Content,
        Author: input.Author,
        ImageURL: input.ImageURL,
    }
    if err := s.repo.Create(report); err != nil {
        return nil, err
    }
    return report, nil
}

func (s *reportService) GetReport(id uuid.UUID) (*Report, error) {
    return s.repo.GetByID(id)
}

func (s *reportService) UpdateReport(id uuid.UUID, input UpdateReportRequest) (*Report, error) {
    report, err := s.repo.GetByID(id)
    if err != nil {
        return nil, err
    }

    if input.Title != "" {
        report.Title = input.Title
    }
    if input.Content != "" {
        report.Content = input.Content
    }
    if input.ImageURL != "" {
        report.ImageURL = input.ImageURL
    }

    if err := s.repo.Update(report); err != nil {
        return nil, err
    }
    return report, nil
}

func (s *reportService) DeleteReport(id uuid.UUID) error {
    return s.repo.Delete(id)
}

func (s *reportService) DeleteAllReports() error {
    return s.repo.DeleteAll()
}

func (s *reportService) ListReports() ([]*Report, error) {
    return s.repo.GetAllTrimmedContent()
}
