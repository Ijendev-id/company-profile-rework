package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"ijendev-backend/internal/middleware"
	"ijendev-backend/internal/models"
)

const postSlugMax = 191

// PostHandler menangani artikel blog (publik dan admin).
type PostHandler struct {
	db *gorm.DB
}

func NewPostHandler(db *gorm.DB) *PostHandler {
	return &PostHandler{db: db}
}

// ---- Bentuk data ----

// postInput adalah isi form artikel dari admin (dipakai untuk buat dan ubah).
type postInput struct {
	Title           string            `json:"title" binding:"required,max=255"`
	Slug            string            `json:"slug" binding:"omitempty,max=191"`
	Excerpt         string            `json:"excerpt" binding:"max=1000"`
	Content         string            `json:"content" binding:"required"`
	CoverImage      string            `json:"cover_image" binding:"max=500"`
	Status          models.PostStatus `json:"status" binding:"required,oneof=draft published"`
	MetaTitle       string            `json:"meta_title" binding:"max=255"`
	MetaDescription string            `json:"meta_description" binding:"max=320"`
	CategoryID      *uint             `json:"category_id"`
}

// normalize merapikan input dan mengembalikan daftar field yang salah (kosong kalau valid).
func (in *postInput) normalize() map[string]string {
	in.Title = strings.TrimSpace(in.Title)
	in.Slug = strings.ToLower(strings.TrimSpace(in.Slug))
	in.Excerpt = strings.TrimSpace(in.Excerpt)
	in.Content = strings.TrimSpace(in.Content)
	in.CoverImage = strings.TrimSpace(in.CoverImage)
	in.MetaTitle = strings.TrimSpace(in.MetaTitle)
	in.MetaDescription = strings.TrimSpace(in.MetaDescription)

	fields := map[string]string{}
	if in.Title == "" {
		fields["title"] = "wajib diisi"
	}
	if in.Content == "" {
		fields["content"] = "wajib diisi"
	}
	if in.Slug != "" && !slugPattern.MatchString(in.Slug) {
		fields["slug"] = "hanya huruf kecil, angka, dan tanda hubung (contoh: judul-artikel)"
	}
	// Hanya gambar dari upload sendiri atau URL http(s); menolak skema berbahaya seperti javascript:.
	if in.CoverImage != "" &&
		!strings.HasPrefix(in.CoverImage, "/uploads/") &&
		!strings.HasPrefix(in.CoverImage, "https://") &&
		!strings.HasPrefix(in.CoverImage, "http://") {
		fields["cover_image"] = "harus berupa hasil upload (/uploads/...) atau URL http(s)"
	}
	return fields
}

// applyTo menyalin input ke model artikel dan mengatur tanggal terbit sesuai status.
func (in *postInput) applyTo(post *models.Post, slug string) {
	post.Title = in.Title
	post.Slug = slug
	post.Excerpt = in.Excerpt
	post.Content = in.Content
	post.CoverImage = in.CoverImage
	post.MetaTitle = in.MetaTitle
	post.MetaDescription = in.MetaDescription
	post.CategoryID = in.CategoryID
	post.Status = in.Status

	if in.Status == models.PostStatusPublished {
		// Tanggal terbit hanya diisi saat pertama kali terbit; edit berikutnya tidak mengubahnya.
		if post.PublishedAt == nil {
			now := time.Now().UTC()
			post.PublishedAt = &now
		}
	} else {
		post.PublishedAt = nil
	}
}

// postSummary adalah bentuk ringkas artikel untuk daftar (tanpa isi lengkap agar respons ringan).
type postSummary struct {
	ID          uint              `json:"id"`
	Title       string            `json:"title"`
	Slug        string            `json:"slug"`
	Excerpt     string            `json:"excerpt"`
	CoverImage  string            `json:"cover_image"`
	Status      models.PostStatus `json:"status"`
	PublishedAt *time.Time        `json:"published_at"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	Category    *models.Category  `json:"category"`
}

// postDetail adalah artikel lengkap untuk halaman baca publik.
type postDetail struct {
	postSummary
	Content         string `json:"content"`
	MetaTitle       string `json:"meta_title"`
	MetaDescription string `json:"meta_description"`
	AuthorName      string `json:"author_name"`
}

// Kolom yang diambil untuk daftar (tanpa `content`).
var summaryColumns = []string{
	"id", "title", "slug", "excerpt", "cover_image", "status",
	"published_at", "created_at", "updated_at", "category_id",
}

func toSummary(p models.Post) postSummary {
	return postSummary{
		ID:          p.ID,
		Title:       p.Title,
		Slug:        p.Slug,
		Excerpt:     p.Excerpt,
		CoverImage:  p.CoverImage,
		Status:      p.Status,
		PublishedAt: p.PublishedAt,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		Category:    p.Category,
	}
}

// publishedOnly membatasi query ke artikel yang sudah terbit dan tanggal terbitnya sudah lewat.
func publishedOnly(db *gorm.DB) *gorm.DB {
	return db.Where("status = ? AND published_at IS NOT NULL AND published_at <= ?",
		models.PostStatusPublished, time.Now().UTC())
}

// ---- Endpoint publik ----

// ListPublished menjawab GET /api/v1/posts?page=1&per_page=10&category=slug&q=kata
func (h *PostHandler) ListPublished(c *gin.Context) {
	page, perPage := parsePagination(c)
	search := strings.TrimSpace(c.Query("q"))
	category := strings.TrimSpace(c.Query("category"))

	// Query dibuat ulang tiap dipakai (Count lalu Find) agar tidak saling memengaruhi.
	build := func() *gorm.DB {
		query := publishedOnly(h.db.WithContext(c.Request.Context()).Model(&models.Post{}))
		if category != "" {
			query = query.Where("category_id IN (?)",
				h.db.Model(&models.Category{}).Select("id").Where("slug = ?", category))
		}
		if search != "" {
			like := likePattern(search)
			query = query.Where("(title LIKE ? OR excerpt LIKE ?)", like, like)
		}
		return query
	}

	h.respondList(c, build, "published_at DESC, id DESC", page, perPage)
}

// GetPublished menjawab GET /api/v1/posts/:slug
func (h *PostHandler) GetPublished(c *gin.Context) {
	var post models.Post
	err := publishedOnly(h.db.WithContext(c.Request.Context())).
		Preload("Category").
		Preload("Author").
		Where("slug = ?", c.Param("slug")).
		First(&post).Error
	if err != nil {
		notFoundOrError(c, err, "artikel tidak ditemukan")
		return
	}

	detail := postDetail{
		postSummary:     toSummary(post),
		Content:         post.Content,
		MetaTitle:       post.MetaTitle,
		MetaDescription: post.MetaDescription,
	}
	if post.Author != nil {
		detail.AuthorName = post.Author.Name
	}
	c.JSON(http.StatusOK, gin.H{"data": detail})
}

// ---- Endpoint admin ----

// ListAdmin menjawab GET /api/v1/admin/posts?status=draft&q=kata&page=1
// Berbeda dengan versi publik, draft juga ikut tampil.
func (h *PostHandler) ListAdmin(c *gin.Context) {
	page, perPage := parsePagination(c)
	search := strings.TrimSpace(c.Query("q"))
	status := strings.TrimSpace(c.Query("status"))

	if status != "" && status != string(models.PostStatusDraft) && status != string(models.PostStatusPublished) {
		invalid(c, map[string]string{"status": "harus draft atau published"})
		return
	}

	build := func() *gorm.DB {
		query := h.db.WithContext(c.Request.Context()).Model(&models.Post{})
		if status != "" {
			query = query.Where("status = ?", status)
		}
		if search != "" {
			like := likePattern(search)
			query = query.Where("(title LIKE ? OR excerpt LIKE ?)", like, like)
		}
		return query
	}

	h.respondList(c, build, "created_at DESC, id DESC", page, perPage)
}

// GetAdmin menjawab GET /api/v1/admin/posts/:id (lengkap, untuk form edit).
func (h *PostHandler) GetAdmin(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}
	h.respondPost(c, id, http.StatusOK)
}

// Create menjawab POST /api/v1/admin/posts.
func (h *PostHandler) Create(c *gin.Context) {
	var in postInput
	if !bindJSON(c, &in) {
		return
	}
	if fields := in.normalize(); len(fields) > 0 {
		invalid(c, fields)
		return
	}
	if !h.categoryExists(c, in.CategoryID) {
		return
	}
	slug, ok := h.resolveSlug(c, in, 0)
	if !ok {
		return
	}

	post := models.Post{AuthorID: middleware.CurrentUserID(c)}
	in.applyTo(&post, slug)

	if err := h.db.WithContext(c.Request.Context()).Create(&post).Error; err != nil {
		writeError(c, err)
		return
	}
	h.respondPost(c, post.ID, http.StatusCreated)
}

// Update menjawab PUT /api/v1/admin/posts/:id (mengganti seluruh isi artikel).
func (h *PostHandler) Update(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	var in postInput
	if !bindJSON(c, &in) {
		return
	}
	if fields := in.normalize(); len(fields) > 0 {
		invalid(c, fields)
		return
	}

	var post models.Post
	if err := h.db.WithContext(c.Request.Context()).First(&post, id).Error; err != nil {
		notFoundOrError(c, err, "artikel tidak ditemukan")
		return
	}
	if !h.categoryExists(c, in.CategoryID) {
		return
	}
	slug, ok := h.resolveSlug(c, in, post.ID)
	if !ok {
		return
	}

	in.applyTo(&post, slug)
	if err := h.db.WithContext(c.Request.Context()).Save(&post).Error; err != nil {
		writeError(c, err)
		return
	}
	h.respondPost(c, post.ID, http.StatusOK)
}

// Delete menjawab DELETE /api/v1/admin/posts/:id.
// Artikel hanya ditandai terhapus (soft delete), datanya masih ada di database.
func (h *PostHandler) Delete(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	result := h.db.WithContext(c.Request.Context()).Unscoped().Delete(&models.Post{}, id)
	if result.Error != nil {
		serverError(c, result.Error)
		return
	}
	if result.RowsAffected == 0 {
		respondError(c, http.StatusNotFound, "artikel tidak ditemukan")
		return
	}
	c.Status(http.StatusNoContent)
}

// ---- Bantuan ----

// respondList menjalankan query daftar (hitung total + ambil satu halaman) dan mengirim hasilnya.
func (h *PostHandler) respondList(c *gin.Context, build func() *gorm.DB, order string, page, perPage int) {
	var total int64
	if err := build().Count(&total).Error; err != nil {
		serverError(c, err)
		return
	}

	var posts []models.Post
	err := build().
		Select(summaryColumns).
		Preload("Category").
		Order(order).
		Offset((page - 1) * perPage).
		Limit(perPage).
		Find(&posts).Error
	if err != nil {
		serverError(c, err)
		return
	}

	items := make([]postSummary, 0, len(posts))
	for _, post := range posts {
		items = append(items, toSummary(post))
	}
	c.JSON(http.StatusOK, gin.H{"data": items, "meta": paginationMeta(page, perPage, total)})
}

// respondPost memuat artikel lengkap (beserta kategori) lalu mengirimnya dengan status HTTP tertentu.
func (h *PostHandler) respondPost(c *gin.Context, id uint, status int) {
	var post models.Post
	if err := h.db.WithContext(c.Request.Context()).Preload("Category").First(&post, id).Error; err != nil {
		notFoundOrError(c, err, "artikel tidak ditemukan")
		return
	}
	c.JSON(status, gin.H{"data": post})
}

// categoryExists memastikan kategori yang dipilih ada. Kategori boleh kosong.
func (h *PostHandler) categoryExists(c *gin.Context, id *uint) bool {
	if id == nil {
		return true
	}

	var count int64
	err := h.db.WithContext(c.Request.Context()).Model(&models.Category{}).Where("id = ?", *id).Count(&count).Error
	if err != nil {
		serverError(c, err)
		return false
	}
	if count == 0 {
		invalid(c, map[string]string{"category_id": "kategori tidak ditemukan"})
		return false
	}
	return true
}

// resolveSlug memakai slug dari input (409 kalau bentrok) atau membuatnya dari judul.
func (h *PostHandler) resolveSlug(c *gin.Context, in postInput, excludeID uint) (string, bool) {
	db := h.db.WithContext(c.Request.Context())

	if in.Slug != "" {
		taken, err := slugTaken(db, &models.Post{}, in.Slug, excludeID)
		if err != nil {
			serverError(c, err)
			return "", false
		}
		if taken {
			respondError(c, http.StatusConflict, "slug sudah dipakai artikel lain")
			return "", false
		}
		return in.Slug, true
	}

	slug, err := uniqueSlug(db, &models.Post{}, slugify(in.Title), postSlugMax, excludeID)
	if err != nil {
		serverError(c, err)
		return "", false
	}
	return slug, true
}
