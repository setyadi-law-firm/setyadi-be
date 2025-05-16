package report

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReportRepository interface {
    Create(report *Report) error
    GetByID(id uuid.UUID) (*Report, error)
    Update(report *Report) error
    Delete(id uuid.UUID) error
    DeleteAll() error
    GetAll() ([]*Report, error)
    GetAllTrimmedContent() ([]*Report, error)
}

type GormReportRepository struct {
    db *gorm.DB
}

func NewGormReportRepository(db *gorm.DB) ReportRepository {
    return &GormReportRepository{db}
}

func (r *GormReportRepository) Create(report *Report) error {
    return r.db.Create(report).Error
}

func (r *GormReportRepository) GetByID(id uuid.UUID) (*Report, error) {
    var report Report
    if err := r.db.First(&report, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &report, nil
}

func (r *GormReportRepository) Update(report *Report) error {
    return r.db.Save(report).Error
}

func (r *GormReportRepository) Delete(id uuid.UUID) error {
    return r.db.Delete(&Report{}, "id = ?", id).Error
}

func (r *GormReportRepository) DeleteAll() error {
    return r.db.Unscoped().Delete(&Report{}).Error
}

func (r *GormReportRepository) GetAll() ([]*Report, error) {
    var reports []*Report
    if err := r.db.Find(&reports).Error; err != nil {
        return nil, err
    }
    return reports, nil
}

func (r *GormReportRepository) GetAllTrimmedContent() ([]*Report, error) {
    var reports []*Report

    fields := []string{
        "id",
        "title",
        fmt.Sprintf("LEFT(content, %d) AS content", MAX_REPORT_CONTENT_PREVIEW_LENGTH),
        "author",
        "image_url",
        "created_at",
        "updated_at",
    }

    err := r.db.
        Model(&Report{}).
        Select(strings.Join(fields, ", ")).
        Find(&reports).Error

    return reports, err
}
