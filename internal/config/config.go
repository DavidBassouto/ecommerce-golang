package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	AWS      AWSConfig
	Upload   UploadConfig
}

type ServerConfig struct {
	Port    string
	GinMode string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type JWTConfig struct {
	SecretKey        string
	ExpiresIn        time.Duration
	RefreshExpiresIn time.Duration
}

type AWSConfig struct {
	Region          string
	AccessKeyID     string
	SecretAccessKey string
	S3BucketName    string
	S3Endpoint      string
}

type UploadConfig struct {
	MaxFileSize int64
	Path        string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	jwtExpiresIn, _ := time.ParseDuration(getEnv("JWT_EXPIRES_IN", "15m"))
	jwtRefreshExpiresIn, _ := time.ParseDuration(getEnv("JWT_REFRESH_EXPIRES_IN", "720h"))
	uploadMaxFileSize, _ := strconv.ParseInt(getEnv("UPLOAD_MAX_FILE_SIZE", "10485760"), 10, 64)

	return &Config{
		Server: ServerConfig{
			Port:    getEnv("SERVER_PORT", "8080"),
			GinMode: getEnv("GIN_MODE", "debug"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "root"),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", "ecommerce"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
		JWT: JWTConfig{
			SecretKey:        getEnv("JWT_SECRET_KEY", "secret"),
			ExpiresIn:        jwtExpiresIn,
			RefreshExpiresIn: jwtRefreshExpiresIn,
		},
		AWS: AWSConfig{
			Region:          getEnv("AWS_REGION", "us-east-1"),
			AccessKeyID:     getEnv("AWS_ACCESS_KEY_ID", "test"),
			SecretAccessKey: getEnv("AWS_SECRET_ACCESS_KEY", "test"),
			S3BucketName:    getEnv("AWS_S3_BUCKET_NAME", "ecommerce-uploads"),
			S3Endpoint:      getEnv("AWS_S3_ENDPOINT", "http://localhost:4566"),
		},
		Upload: UploadConfig{
			MaxFileSize: uploadMaxFileSize,
			Path:        getEnv("UPLOAD_PATH", "./uploads"),
		},
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
