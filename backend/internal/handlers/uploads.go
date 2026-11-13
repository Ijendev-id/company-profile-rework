package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// UploadHandler menangani upload gambar (cover artikel).
type UploadHandler struct {
	dir      string
	maxBytes int64
}

func NewUploadHandler(dir string, maxMB int) *UploadHandler {
	return &UploadHandler{dir: dir, maxBytes: int64(maxMB) << 20}
}

// Hanya gambar raster. SVG sengaja tidak diizinkan karena bisa memuat script.
var allowedImageTypes = map[string]string{
	"image/jpeg": ".jpg",
	"image/png":  ".png",
	"image/gif":  ".gif",
	"image/webp": ".webp",
}

// Image menjawab POST /api/v1/admin/uploads (multipart/form-data, field "file").
// Jenis file ditentukan dari isinya, bukan dari nama atau ekstensi yang dikirim client.
func (h *UploadHandler) Image(c *gin.Context) {
	// Batasi ukuran body sejak awal agar file raksasa tidak dibaca sampai habis.
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, h.maxBytes+(1<<20))

	header, err := c.FormFile("file")
	if err != nil {
		var tooBig *http.MaxBytesError
		if errors.As(err, &tooBig) {
			respondError(c, http.StatusRequestEntityTooLarge, fmt.Sprintf("ukuran file maksimal %d MB", h.maxBytes>>20))
			return
		}
		respondError(c, http.StatusBadRequest, `kirim file gambar dengan nama field "file"`)
		return
	}
	if header.Size > h.maxBytes {
		respondError(c, http.StatusRequestEntityTooLarge, fmt.Sprintf("ukuran file maksimal %d MB", h.maxBytes>>20))
		return
	}

	src, err := header.Open()
	if err != nil {
		serverError(c, err)
		return
	}
	defer src.Close()

	head := make([]byte, 512)
	n, err := io.ReadFull(src, head)
	if err != nil && !errors.Is(err, io.ErrUnexpectedEOF) && !errors.Is(err, io.EOF) {
		serverError(c, err)
		return
	}
	ext, ok := allowedImageTypes[http.DetectContentType(head[:n])]
	if !ok {
		respondError(c, http.StatusUnsupportedMediaType, "file harus berupa gambar JPG, PNG, GIF, atau WebP")
		return
	}
	if _, err := src.Seek(0, io.SeekStart); err != nil {
		serverError(c, err)
		return
	}

	name, err := randomName()
	if err != nil {
		serverError(c, err)
		return
	}
	filename := name + ext

	dst, err := os.OpenFile(filepath.Join(h.dir, filename), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
	if err != nil {
		serverError(c, err)
		return
	}
	if _, err := io.Copy(dst, src); err != nil {
		dst.Close()
		_ = os.Remove(dst.Name())
		serverError(c, err)
		return
	}
	if err := dst.Close(); err != nil {
		log.Printf("upload: gagal menutup file %s: %v", filename, err)
		serverError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"url": "/uploads/" + filename})
}

// randomName membuat nama file acak (nama asli dari client tidak pernah dipakai).
func randomName() (string, error) {
	buf := make([]byte, 16)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return hex.EncodeToString(buf), nil
}
