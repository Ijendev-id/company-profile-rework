// Package handlers berisi fungsi yang menangani request HTTP (pengganti Controller di Laravel).
package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HealthHandler mengecek apakah server dan database sehat.
type HealthHandler struct {
	db *gorm.DB
}

func NewHealthHandler(db *gorm.DB) *HealthHandler {
	return &HealthHandler{db: db}
}

// Check menjawab GET /health.
func (h *HealthHandler) Check(c *gin.Context) {
	sqlDB, err := h.db.DB()
	if err == nil {
		err = sqlDB.PingContext(c.Request.Context())
	}
	if err != nil {
		// Detail error hanya di log server, tidak dikirim ke client.
		log.Printf("health check: database tidak sehat: %v", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": "error", "database": "down"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "database": "up"})
}
