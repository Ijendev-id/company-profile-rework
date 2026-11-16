// Package media membantu mengelola file hasil upload (folder uploads/).
// Hanya memakai library standar agar mudah diuji.
package media

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
)

// Nama file upload selalu 32 karakter hex + ekstensi gambar (lihat handlers/uploads.go).
// Karakter sebelum "/uploads/" tidak boleh huruf, angka, titik, atau "-" supaya alamat milik
// situs lain (https://situs-lain.com/uploads/...) tidak ikut terhitung sebagai file kita.
var uploadRef = regexp.MustCompile(`(?:^|[^A-Za-z0-9.\-])/uploads/([a-f0-9]{32}\.(?:jpg|png|gif|webp))`)

// ExtractNames mengambil nama file upload yang disebut dalam teks-teks (gambar sampul, isi Markdown).
// Hasilnya unik, urutan sesuai kemunculan pertama.
func ExtractNames(texts ...string) []string {
	var names []string
	seen := map[string]bool{}
	for _, text := range texts {
		for _, m := range uploadRef.FindAllStringSubmatch(text, -1) {
			if !seen[m[1]] {
				seen[m[1]] = true
				names = append(names, m[1])
			}
		}
	}
	return names
}

// Remove menghapus satu file upload dari folder dir. Nama harus berbentuk nama upload yang sah
// (bukan path), dan file yang memang sudah tidak ada dianggap berhasil.
func Remove(dir, name string) error {
	if !uploadRef.MatchString("/uploads/"+name) || filepath.Base(name) != name {
		return errors.New("nama file upload tidak valid")
	}
	if err := os.Remove(filepath.Join(dir, name)); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	return nil
}
