package handlers

import (
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func init() {
	// Nama field di pesan error memakai nama JSON (misalnya "meta_title"), bukan nama field Go.
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(f reflect.StructField) string {
			name := strings.SplitN(f.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}

// ---- Response error ----

func respondError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

// serverError mencatat detail error di log server, tetapi hanya mengirim pesan umum ke client.
func serverError(c *gin.Context, err error) {
	log.Printf("%s %s: %v", c.Request.Method, c.Request.URL.Path, err)
	respondError(c, http.StatusInternalServerError, "terjadi kesalahan pada server")
}

func invalid(c *gin.Context, fields map[string]string) {
	c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "data tidak valid", "fields": fields})
}

func notFoundOrError(c *gin.Context, err error, message string) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		respondError(c, http.StatusNotFound, message)
		return
	}
	serverError(c, err)
}

// writeError menerjemahkan error database: nilai unik yang bentrok menjadi 409, selebihnya 500.
func writeError(c *gin.Context, err error) {
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		respondError(c, http.StatusConflict, "data dengan nilai yang sama sudah ada")
		return
	}
	serverError(c, err)
}

// ---- Input ----

// bindJSON membaca body JSON dan menjalankan aturan validasi `binding:"..."`.
// Mengembalikan false (dan sudah mengirim response error) kalau gagal.
func bindJSON(c *gin.Context, dst any) bool {
	err := c.ShouldBindJSON(dst)
	if err == nil {
		return true
	}

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		fields := make(map[string]string, len(validationErrors))
		for _, fe := range validationErrors {
			fields[fe.Field()] = validationMessage(fe)
		}
		invalid(c, fields)
		return false
	}

	respondError(c, http.StatusBadRequest, "format JSON tidak valid")
	return false
}

func validationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "wajib diisi"
	case "max":
		return fmt.Sprintf("terlalu panjang (maksimal %s)", fe.Param())
	case "min":
		return fmt.Sprintf("terlalu pendek (minimal %s)", fe.Param())
	case "email":
		return "format email tidak valid"
	case "oneof":
		return "nilai tidak valid"
	default:
		return "tidak valid"
	}
}

// parseID membaca parameter URL :id.
func parseID(c *gin.Context) (uint, bool) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil || id == 0 {
		respondError(c, http.StatusBadRequest, "id tidak valid")
		return 0, false
	}
	return uint(id), true
}

// ---- Pagination & pencarian ----

const (
	defaultPerPage = 10
	maxPerPage     = 50
)

func parsePagination(c *gin.Context) (page, perPage int) {
	page = atoiDefault(c.Query("page"), 1)
	if page < 1 {
		page = 1
	}
	perPage = atoiDefault(c.Query("per_page"), defaultPerPage)
	if perPage < 1 {
		perPage = defaultPerPage
	}
	if perPage > maxPerPage {
		perPage = maxPerPage
	}
	return page, perPage
}

func paginationMeta(page, perPage int, total int64) gin.H {
	return gin.H{
		"page":        page,
		"per_page":    perPage,
		"total":       total,
		"total_pages": int(math.Ceil(float64(total) / float64(perPage))),
	}
}

func atoiDefault(value string, fallback int) int {
	n, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return n
}

var likeEscaper = strings.NewReplacer(`\`, `\\`, `%`, `\%`, `_`, `\_`)

// likePattern menyiapkan kata pencarian untuk LIKE '%kata%' (karakter % dan _ dari pengguna dibuat harfiah).
func likePattern(search string) string {
	return "%" + likeEscaper.Replace(search) + "%"
}

// ---- Slug ----

var (
	nonSlugChars = regexp.MustCompile(`[^a-z0-9]+`)
	slugPattern  = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)
)

// slugify mengubah teks menjadi slug URL: huruf kecil, angka, dan tanda hubung.
// Huruf non-ASCII (misalnya beraksen) dibuang. Kalau hasilnya kosong, dipakai "post".
func slugify(text string) string {
	slug := nonSlugChars.ReplaceAllString(strings.ToLower(strings.TrimSpace(text)), "-")
	slug = strings.Trim(slug, "-")
	if len(slug) > 150 {
		slug = strings.Trim(slug[:150], "-")
	}
	if slug == "" {
		return "post"
	}
	return slug
}

// slugTaken memeriksa apakah slug sudah dipakai (termasuk oleh data yang sudah di-soft-delete).
func slugTaken(db *gorm.DB, model any, slug string, excludeID uint) (bool, error) {
	var count int64
	err := db.Unscoped().Model(model).Where("slug = ? AND id <> ?", slug, excludeID).Count(&count).Error
	return count > 0, err
}

// uniqueSlug menambahkan akhiran -2, -3, dst. sampai slug belum dipakai.
func uniqueSlug(db *gorm.DB, model any, base string, maxLen int, excludeID uint) (string, error) {
	if limit := maxLen - 8; len(base) > limit {
		base = strings.Trim(base[:limit], "-")
	}

	candidate := base
	for i := 2; i <= 1000; i++ {
		taken, err := slugTaken(db, model, candidate, excludeID)
		if err != nil {
			return "", err
		}
		if !taken {
			return candidate, nil
		}
		candidate = fmt.Sprintf("%s-%d", base, i)
	}
	return "", errors.New("gagal membuat slug unik")
}
