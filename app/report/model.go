package report

import (
    "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

type Report struct {
    ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    AuthorID  uint      `json:"author_id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

func (r *Report) BeforeCreate(tx *gorm.DB) (err error) {
    if r.ID == uuid.Nil {
        r.ID = uuid.New()
    }
    return
}
