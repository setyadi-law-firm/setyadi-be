package image

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"
)

type ImageService interface {
	UploadImage(file *multipart.FileHeader) (string, error)
}

type ImageServiceImpl struct {
	cloudName    string
	apiKey    string
	apiSecret    string
}

func NewImageService(cloudName, apiKey, apiSecret string) ImageService {
	return &ImageServiceImpl{
		cloudName: cloudName,
		apiKey:    apiKey,
		apiSecret: apiSecret,
	}
}

func (s *ImageServiceImpl) UploadImage(file *multipart.FileHeader) (string, error) {
	uploadURL := fmt.Sprintf("https://api.cloudinary.com/v1_1/%s/image/upload", s.cloudName)

	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer src.Close()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	part, err := writer.CreateFormFile("file", file.Filename)
	if err != nil {
		return "", fmt.Errorf("create form file failed: %w", err)
	}
	if _, err := io.Copy(part, src); err != nil {
		return "", fmt.Errorf("copy file failed: %w", err)
	}

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	toSign := "timestamp=" + timestamp + s.apiSecret
	h := sha1.New()
	h.Write([]byte(toSign))
	signature := fmt.Sprintf("%x", h.Sum(nil))

	_ = writer.WriteField("api_key", s.apiKey)
	_ = writer.WriteField("timestamp", timestamp)
	_ = writer.WriteField("signature", signature)

	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("multipart close failed: %w", err)
	}

	req, err := http.NewRequest("POST", uploadURL, &body)
	if err != nil {
		return "", fmt.Errorf("create request failed: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("upload failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("upload failed: %s", b)
	}

	var result struct {
		SecureURL string `json:"secure_url"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("decode failed: %w", err)
	}

	return result.SecureURL, nil
}
