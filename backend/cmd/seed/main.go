// Command seed membuat akun admin (atau mengganti passwordnya dengan -reset).
//
// Cara pakai:
//
//	go run ./cmd/seed -name "Nama Admin" -email admin@contoh.com
//	go run ./cmd/seed -email admin@contoh.com -reset      (ganti password)
//
// Password diketik lewat prompt (tidak tampil di layar dan tidak tersimpan di riwayat terminal).
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/term"

	"ijendev-backend/internal/auth"
	"ijendev-backend/internal/config"
	"ijendev-backend/internal/database"
	"ijendev-backend/internal/models"
)

const minPasswordLength = 10

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	name := flag.String("name", "Admin", "nama admin")
	email := flag.String("email", "", "email admin (dipakai untuk login)")
	reset := flag.Bool("reset", false, "ganti password admin yang sudah ada")
	flag.Parse()

	*email = strings.ToLower(strings.TrimSpace(*email))
	*name = strings.TrimSpace(*name)
	if !strings.Contains(*email, "@") {
		return fmt.Errorf("isi email dengan benar, contoh: go run ./cmd/seed -email admin@contoh.com")
	}
	if *name == "" {
		return fmt.Errorf("nama tidak boleh kosong")
	}

	password, err := promptPassword()
	if err != nil {
		return err
	}
	hash, err := auth.HashPassword(password)
	if err != nil {
		return fmt.Errorf("membuat hash password: %w", err)
	}

	cfg := config.Load()
	db, err := database.Connect(cfg)
	if err != nil {
		return fmt.Errorf("koneksi database gagal: %w", err)
	}
	if err := database.Migrate(db); err != nil {
		return fmt.Errorf("migrasi tabel gagal: %w", err)
	}

	var existing models.User
	result := db.Where("email = ?", *email).Limit(1).Find(&existing)
	if result.Error != nil {
		return fmt.Errorf("mencari user: %w", result.Error)
	}

	if result.RowsAffected > 0 {
		if !*reset {
			return fmt.Errorf("email %s sudah terdaftar. Pakai -reset kalau ingin mengganti passwordnya", *email)
		}
		if err := db.Model(&existing).Update("password_hash", hash).Error; err != nil {
			return fmt.Errorf("mengganti password: %w", err)
		}
		fmt.Println("Password admin berhasil diganti.")
		return nil
	}

	if *reset {
		return fmt.Errorf("email %s belum terdaftar, jalankan tanpa -reset untuk membuatnya", *email)
	}

	user := models.User{Name: *name, Email: *email, PasswordHash: hash}
	if err := db.Create(&user).Error; err != nil {
		return fmt.Errorf("membuat admin: %w", err)
	}
	fmt.Printf("Admin berhasil dibuat: %s <%s>\n", user.Name, user.Email)
	return nil
}

func promptPassword() (string, error) {
	fd := int(os.Stdin.Fd())
	if !term.IsTerminal(fd) {
		return "", fmt.Errorf("jalankan perintah ini langsung di terminal (bukan lewat pipe) agar password bisa diketik")
	}

	fmt.Printf("Password (minimal %d karakter): ", minPasswordLength)
	first, err := term.ReadPassword(fd)
	fmt.Println()
	if err != nil {
		return "", fmt.Errorf("membaca password: %w", err)
	}

	fmt.Print("Ulangi password: ")
	second, err := term.ReadPassword(fd)
	fmt.Println()
	if err != nil {
		return "", fmt.Errorf("membaca password: %w", err)
	}

	password := string(first)
	if password != string(second) {
		return "", fmt.Errorf("kedua password tidak sama")
	}
	if len(password) < minPasswordLength {
		return "", fmt.Errorf("password minimal %d karakter", minPasswordLength)
	}
	if len(password) > auth.MaxPasswordBytes {
		return "", fmt.Errorf("password maksimal %d karakter", auth.MaxPasswordBytes)
	}
	return password, nil
}
