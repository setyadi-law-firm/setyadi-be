package models

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	JwtSecretKey           string
	JwtExpiryInDays        int
	JwtRefreshExpiryInDays int

	SupabaseURL        string
	SupabaseAPIKey     string
	SupabaseBucketName string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	jwtExpiry, err := strconv.Atoi(getEnv("JWT_EXPIRY_DAYS", "1"))
	if err != nil {
		log.Fatal("invalid JWT_EXPIRY_DAYS:", err)
	}

	refreshExpiry, err := strconv.Atoi(getEnv("JWT_REFRESH_EXPIRY_DAYS", "7"))
	if err != nil {
		log.Fatal("invalid JWT_REFRESH_EXPIRY_DAYS:", err)
	}

	return &Config{
		JwtSecretKey:           getEnv("JWT_SECRET_KEY", "default-secret"),
		JwtExpiryInDays:        jwtExpiry,
		JwtRefreshExpiryInDays: refreshExpiry,

		SupabaseURL:        getEnv("SUPABASE_URL", ""),
		SupabaseAPIKey:     getEnv("SUPABASE_API_KEY", ""),
		SupabaseBucketName: getEnv("SUPABASE_BUCKET_NAME", ""),
		DBHost:             getEnv("DB_HOST", "localhost"),
		DBPort:             getEnv("DB_PORT", "5432"),
		DBUser:             getEnv("DB_USER", "postgres"),
		DBPassword:         getEnv("DB_PASSWORD", ""),
		DBName:             getEnv("DB_NAME", "authdb"),
		DBSSLMode:          getEnv("DB_SSL_MODE", "disable"),
	}
}

func (c *Config) Dsn() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta",
		c.DBHost, c.DBUser, c.DBPassword, c.DBName, c.DBPort, c.DBSSLMode,
	)
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
