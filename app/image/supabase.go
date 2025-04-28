package image

import (
    "bytes"
    "fmt"
    "io"
    "mime/multipart"
    "net/http"
    "github.com/setyadi-law-firm/setyadi-be/app/models" 
)

type Supabase struct {
	config *models.Config
}

func NewSupabase(config *models.Config) *Supabase {
	return &Supabase{config}
}

func (s *Supabase) UploadToSupabase(file *multipart.FileHeader) (string, error) {
    src, err := file.Open()
    if err != nil {
        return "", err
    }
    defer src.Close()

    fileBytes, err := io.ReadAll(src)
    if err != nil {
        return "", err
    }

    uploadURL := fmt.Sprintf("%s/storage/v1/object/%s/%s", 
        s.config.SupabaseURL, 
        s.config.SupabaseBucketName, 
        file.Filename,
    )

    // Buat request POST
    req, err := http.NewRequest("POST", uploadURL, bytes.NewReader(fileBytes))
    if err != nil {
        return "", err
    }

    req.Header.Set("Authorization", "Bearer " + s.config.SupabaseAPIKey)
    req.Header.Set("Content-Type", file.Header.Get("Content-Type"))
    req.Header.Set("x-upsert", "true")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode >= 300 {
        bodyBytes, _ := io.ReadAll(resp.Body)
        return "", fmt.Errorf("failed upload: %s", string(bodyBytes))
    }

    // Success
    publicURL := fmt.Sprintf("%s/storage/v1/object/public/%s/%s",
        s.config.SupabaseURL,
        s.config.SupabaseBucketName,
        file.Filename,
    )
    return publicURL, nil
}
