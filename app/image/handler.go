package image

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type ImageHandler struct {
    service ImageService
}

func NewImageHandler(service ImageService) *ImageHandler {
    return &ImageHandler{service: service}
}

func (h *ImageHandler) UploadImage(c *gin.Context) {
    var req UploadImageRequest
    if err := c.ShouldBind(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "error":   err.Error(),
            "data":    nil,
        })
        return
    }

    url, err := h.service.UploadImage(req.File)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "success": false,
            "error":   err.Error(),
            "data":    nil,
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "error":   nil,
        "data": UploadImageResponse{
            URL: url,
        },
    })
}
