package image

import (
    "mime/multipart"
)

type ImageService interface {
    UploadImage(file *multipart.FileHeader) (string, error)
}

type ImageServiceImpl struct {
    supabase *Supabase
}

func NewImageService(supabase *Supabase) ImageService {
    return &ImageServiceImpl{supabase: supabase}
}

func (s *ImageServiceImpl) UploadImage(file *multipart.FileHeader) (string, error) {
    return s.supabase.UploadToSupabase(file)
}
