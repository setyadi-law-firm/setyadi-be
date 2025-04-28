package image

import "mime/multipart"

type UploadImageRequest struct {
    File *multipart.FileHeader `form:"file" binding:"required"`
}

type UploadImageResponse struct {
    URL string `json:"url"`
}
