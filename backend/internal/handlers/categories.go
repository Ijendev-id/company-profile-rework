package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"ijendev-backend/internal/models"
)

const categorySlugMax = 120

// CategoryHandler menangani kategori artikel.
type CategoryHandler struct {
	db *gorm.DB
}

func NewCategoryHandler(db *gorm.DB) *CategoryHandler {
	return &CategoryHandler{db: db}
}

type categoryInput struct {
	Name string `json:"name" binding:"required,max=100"`
	Slug string `json:"slug" binding:"omitempty,max=120"`
}

// normalize merapikan input dan mengembalikan daftar field yang salah (kosong kalau valid).
func (in *categoryInput) normalize() map[string]string {
	in.Name = strings.TrimSpace(in.Name)
	in.Slug = strings.ToLower(strings.TrimSpace(in.Slug))

	fields := map[string]string{}
	if in.Name == "" {
		fields["name"] = "wajib diisi"
	}
	if in.Slug != "" && !slugPattern.MatchString(in.Slug) {
		fields["slug"] = "hanya huruf kecil, angka, dan tanda hubung (contoh: berita-terbaru)"
	}
	return fields
}

// List menjawab GET /api/v1/categories (publik).
func (h *CategoryHandler) List(c *gin.Context) {
	var categories []models.Category
	if err := h.db.WithContext(c.Request.Context()).Order("name ASC").Find(&categories).Error; err != nil {
		serverError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// Create menjawab POST /api/v1/admin/categories.
func (h *CategoryHandler) Create(c *gin.Context) {
	var in categoryInput
	if !bindJSON(c, &in) {
		return
	}
	if fields := in.normalize(); len(fields) > 0 {
		invalid(c, fields)
		return
	}

	slug, ok := h.resolveSlug(c, in, 0)
	if !ok {
		return
	}

	category := models.Category{Name: in.Name, Slug: slug}
	if err := h.db.WithContext(c.Request.Context()).Create(&category).Error; err != nil {
		writeError(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": category})
}

// Update menjawab PUT /api/v1/admin/categories/:id.
func (h *CategoryHandler) Update(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	var in categoryInput
	if !bindJSON(c, &in) {
		return
	}
	if fields := in.normalize(); len(fields) > 0 {
		invalid(c, fields)
		return
	}

	var category models.Category
	if err := h.db.WithContext(c.Request.Context()).First(&category, id).Error; err != nil {
		notFoundOrError(c, err, "kategori tidak ditemukan")
		return
	}

	slug, ok := h.resolveSlug(c, in, category.ID)
	if !ok {
		return
	}

	category.Name = in.Name
	category.Slug = slug
	if err := h.db.WithContext(c.Request.Context()).Save(&category).Error; err != nil {
		writeError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": category})
}

// Delete menjawab DELETE /api/v1/admin/categories/:id.
// Artikel di kategori ini tidak ikut terhapus; kategorinya menjadi kosong.
func (h *CategoryHandler) Delete(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	result := h.db.WithContext(c.Request.Context()).Delete(&models.Category{}, id)
	if result.Error != nil {
		serverError(c, result.Error)
		return
	}
	if result.RowsAffected == 0 {
		respondError(c, http.StatusNotFound, "kategori tidak ditemukan")
		return
	}
	c.Status(http.StatusNoContent)
}

// resolveSlug memakai slug dari input (409 kalau bentrok) atau membuatnya dari nama.
func (h *CategoryHandler) resolveSlug(c *gin.Context, in categoryInput, excludeID uint) (string, bool) {
	db := h.db.WithContext(c.Request.Context())

	if in.Slug != "" {
		taken, err := slugTaken(db, &models.Category{}, in.Slug, excludeID)
		if err != nil {
			serverError(c, err)
			return "", false
		}
		if taken {
			respondError(c, http.StatusConflict, "slug sudah dipakai kategori lain")
			return "", false
		}
		return in.Slug, true
	}

	slug, err := uniqueSlug(db, &models.Category{}, slugify(in.Name), categorySlugMax, excludeID)
	if err != nil {
		serverError(c, err)
		return "", false
	}
	return slug, true
}
