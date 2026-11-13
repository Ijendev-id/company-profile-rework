package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"ijendev-backend/internal/auth"
	"ijendev-backend/internal/middleware"
	"ijendev-backend/internal/models"
)

// AuthHandler menangani login admin.
type AuthHandler struct {
	db     *gorm.DB
	tokens *auth.TokenManager
}

func NewAuthHandler(db *gorm.DB, tokens *auth.TokenManager) *AuthHandler {
	return &AuthHandler{db: db, tokens: tokens}
}

type loginInput struct {
	Email    string `json:"email" binding:"required,email,max=191"`
	Password string `json:"password" binding:"required,max=72"`
}

// Login menjawab POST /api/v1/auth/login dan mengembalikan token JWT.
func (h *AuthHandler) Login(c *gin.Context) {
	var in loginInput
	if !bindJSON(c, &in) {
		return
	}
	email := strings.ToLower(strings.TrimSpace(in.Email))

	const failMessage = "email atau password salah"

	var user models.User
	err := h.db.WithContext(c.Request.Context()).Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		auth.FakeCheck(in.Password) // samakan waktu respons dengan email yang terdaftar
		respondError(c, http.StatusUnauthorized, failMessage)
		return
	}
	if err != nil {
		serverError(c, err)
		return
	}

	if !auth.CheckPassword(user.PasswordHash, in.Password) {
		respondError(c, http.StatusUnauthorized, failMessage)
		return
	}

	token, expiresAt, err := h.tokens.Generate(user.ID)
	if err != nil {
		serverError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":      token,
		"expires_at": expiresAt,
		"user":       user,
	})
}

// Me menjawab GET /api/v1/admin/me: data admin yang sedang login (dipakai frontend untuk cek sesi).
func (h *AuthHandler) Me(c *gin.Context) {
	var user models.User
	err := h.db.WithContext(c.Request.Context()).First(&user, middleware.CurrentUserID(c)).Error
	if err != nil {
		notFoundOrError(c, err, "user tidak ditemukan")
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
