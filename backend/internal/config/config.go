// Package config membaca pengaturan aplikasi dari file .env / environment.
package config

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// Config menampung seluruh pengaturan aplikasi.
type Config struct {
	AppEnv  string
	AppPort string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	AllowedOrigins []string
	AutoMigrate    bool

	// Login admin (JWT)
	JWTSecret      string
	JWTExpiryHours int

	// Upload gambar
	UploadDir   string
	MaxUploadMB int
}

// Load membaca file .env (kalau ada) lalu mengisi Config.
// Di server produksi, nilai biasanya langsung dari environment sehingga .env tidak wajib ada.
func Load() *Config {
	_ = godotenv.Load()

	return &Config{
		AppEnv:  getEnv("APP_ENV", "development"),
		AppPort: getEnv("APP_PORT", "8080"),

		DBHost:     getEnv("DB_HOST", "127.0.0.1"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "new-ijendev"),

		AllowedOrigins: splitList(getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:3000")),
		AutoMigrate:    getEnvBool("AUTO_MIGRATE", true),

		JWTSecret:      getEnv("JWT_SECRET", ""),
		JWTExpiryHours: getEnvInt("JWT_EXPIRY_HOURS", 24),

		UploadDir:   getEnv("UPLOAD_DIR", "uploads"),
		MaxUploadMB: getEnvInt("MAX_UPLOAD_MB", 5),
	}
}

// Validate memastikan pengaturan wajib sudah benar sebelum server dijalankan.
func (c *Config) Validate() error {
	if len(c.JWTSecret) < 32 {
		return errors.New(`JWT_SECRET di file .env belum diisi atau kurang dari 32 karakter. ` +
			`Buat dengan perintah: node -e "console.log(require('crypto').randomBytes(48).toString('base64url'))" ` +
			`lalu salin hasilnya ke JWT_SECRET=`)
	}
	if c.JWTExpiryHours < 1 {
		return errors.New("JWT_EXPIRY_HOURS minimal 1")
	}
	if c.MaxUploadMB < 1 {
		return errors.New("MAX_UPLOAD_MB minimal 1")
	}
	return nil
}

// IsProduction true kalau APP_ENV=production.
func (c *Config) IsProduction() bool {
	return c.AppEnv == "production"
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}

func getEnvBool(key string, fallback bool) bool {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		return fallback
	}
	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return fallback
	}
	return parsed
}

func getEnvInt(key string, fallback int) int {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		return fallback
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return parsed
}

func splitList(value string) []string {
	var items []string
	for _, part := range strings.Split(value, ",") {
		if part = strings.TrimSpace(part); part != "" {
			items = append(items, part)
		}
	}
	return items
}
