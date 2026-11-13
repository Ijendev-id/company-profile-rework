package models

import (
	"time"

	"gorm.io/gorm"
)

// PostStatus adalah status terbit artikel.
type PostStatus string

const (
	PostStatusDraft     PostStatus = "draft"
	PostStatusPublished PostStatus = "published"
)

// Post adalah artikel blog.
type Post struct {
	ID uint `gorm:"primaryKey" json:"id"`

	Title string `gorm:"size:255;not null" json:"title"`
	// Slug dipakai di URL, contoh /blog/judul-artikel. Panjang 191 agar aman untuk index unik utf8mb4.
	Slug    string `gorm:"size:191;uniqueIndex;not null" json:"slug"`
	Excerpt string `gorm:"type:text" json:"excerpt"`
	// Content disimpan sebagai Markdown.
	Content    string `gorm:"type:longtext;not null" json:"content"`
	CoverImage string `gorm:"size:500" json:"cover_image"`

	Status      PostStatus `gorm:"size:20;not null;default:draft;index" json:"status"`
	PublishedAt *time.Time `gorm:"index" json:"published_at"`

	// SEO opsional. Kalau kosong, frontend memakai Title dan Excerpt.
	MetaTitle       string `gorm:"size:255" json:"meta_title"`
	MetaDescription string `gorm:"size:320" json:"meta_description"`

	// Kategori boleh kosong; jika kategori dihapus, artikel tetap ada (category_id jadi NULL).
	CategoryID *uint     `gorm:"index" json:"category_id"`
	Category   *Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"category,omitempty"`

	// Penulis tidak ikut ke JSON agar email admin tidak bocor di API publik.
	AuthorID uint  `gorm:"not null;index" json:"author_id"`
	Author   *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"-"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// Soft delete: artikel yang "dihapus" hanya ditandai, datanya masih ada di database.
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
