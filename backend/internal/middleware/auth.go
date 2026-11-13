package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"ijendev-backend/internal/auth"
	"ijendev-backend/internal/models"
)

const ctxUserID = "userID"

// RequireAuth hanya meloloskan request dengan header "Authorization: Bearer <token>" yang valid
// dan milik user yang masih ada di database.
func RequireAuth(tokens *auth.TokenManager, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		const prefix = "Bearer "
		header := c.GetHeader("Authorization")
		if len(header) <= len(prefix) || !strings.EqualFold(header[:len(prefix)], prefix) {
			unauthorized(c)
			return
		}

		userID, err := tokens.Parse(strings.TrimSpace(header[len(prefix):]))
		if err != nil {
			unauthorized(c)
			return
		}

		var count int64
		err = db.WithContext(c.Request.Context()).Model(&models.User{}).Where("id = ?", userID).Count(&count).Error
		if err != nil {
			log.Printf("auth: gagal memeriksa user: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "terjadi kesalahan pada server"})
			return
		}
		if count == 0 {
			unauthorized(c)
			return
		}

		c.Set(ctxUserID, userID)
		c.Next()
	}
}

// CurrentUserID mengambil ID user yang sedang login (hanya terisi setelah RequireAuth).
func CurrentUserID(c *gin.Context) uint {
	value, _ := c.Get(ctxUserID)
	id, _ := value.(uint)
	return id
}

func unauthorized(c *gin.Context) {
	c.Header("WWW-Authenticate", "Bearer")
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "silakan login terlebih dahulu"})
}
