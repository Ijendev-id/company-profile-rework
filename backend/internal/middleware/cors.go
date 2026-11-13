// Package middleware berisi kode yang berjalan sebelum/sesudah handler (mirip middleware Laravel).
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORS mengizinkan frontend (misalnya Nuxt di localhost:3000) memanggil API ini dari browser.
// Hanya origin yang terdaftar di CORS_ALLOWED_ORIGINS yang diizinkan.
//
// Catatan: kalau nanti login memakai cookie, tambahkan header
// "Access-Control-Allow-Credentials: true". Kalau memakai header Authorization (token), tidak perlu.
func CORS(allowedOrigins []string) gin.HandlerFunc {
	allowed := make(map[string]struct{}, len(allowedOrigins))
	for _, origin := range allowedOrigins {
		allowed[origin] = struct{}{}
	}

	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if _, ok := allowed[origin]; ok {
			header := c.Writer.Header()
			header.Set("Access-Control-Allow-Origin", origin)
			header.Set("Vary", "Origin")
			header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			header.Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Accept")
			header.Set("Access-Control-Max-Age", "600")
		}

		// Request "preflight" dari browser cukup dijawab tanpa isi.
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
