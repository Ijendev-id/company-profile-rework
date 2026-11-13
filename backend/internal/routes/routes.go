// Package routes mendaftarkan semua URL API (pengganti routes/web.php).
package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"ijendev-backend/internal/auth"
	"ijendev-backend/internal/config"
	"ijendev-backend/internal/handlers"
	"ijendev-backend/internal/middleware"
)

// Setup membuat router Gin lengkap dengan middleware dan daftar route.
func Setup(cfg *config.Config, db *gorm.DB) *gin.Engine {
	if cfg.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	// Tidak mempercayai header proxy (X-Forwarded-For). Kalau nanti di belakang Nginx, atur daftar proxy di sini.
	_ = router.SetTrustedProxies(nil)
	router.Use(gin.Logger(), gin.Recovery(), middleware.CORS(cfg.AllowedOrigins))

	tokens := auth.NewTokenManager(cfg.JWTSecret, time.Duration(cfg.JWTExpiryHours)*time.Hour)

	authHandler := handlers.NewAuthHandler(db, tokens)
	postHandler := handlers.NewPostHandler(db)
	categoryHandler := handlers.NewCategoryHandler(db)
	uploadHandler := handlers.NewUploadHandler(cfg.UploadDir, cfg.MaxUploadMB)

	router.GET("/health", handlers.NewHealthHandler(db).Check)

	// Gambar upload disajikan sebagai file statis (daftar isi folder tidak ditampilkan).
	uploads := router.Group("/uploads")
	uploads.Use(func(c *gin.Context) {
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Cache-Control", "public, max-age=31536000, immutable")
		c.Next()
	})
	uploads.Static("/", cfg.UploadDir)

	v1 := router.Group("/api/v1")

	// Publik (tanpa login)
	v1.POST("/auth/login", middleware.RateLimit(10, time.Minute), authHandler.Login)
	v1.GET("/posts", postHandler.ListPublished)
	v1.GET("/posts/:slug", postHandler.GetPublished)
	v1.GET("/categories", categoryHandler.List)

	// Admin (wajib login)
	admin := v1.Group("/admin", middleware.RequireAuth(tokens, db))
	admin.GET("/me", authHandler.Me)

	admin.GET("/posts", postHandler.ListAdmin)
	admin.GET("/posts/:id", postHandler.GetAdmin)
	admin.POST("/posts", postHandler.Create)
	admin.PUT("/posts/:id", postHandler.Update)
	admin.DELETE("/posts/:id", postHandler.Delete)

	admin.POST("/categories", categoryHandler.Create)
	admin.PUT("/categories/:id", categoryHandler.Update)
	admin.DELETE("/categories/:id", categoryHandler.Delete)

	admin.POST("/uploads", uploadHandler.Image)

	return router
}
