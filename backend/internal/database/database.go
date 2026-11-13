// Package database mengurus koneksi MySQL dan pembuatan tabel.
package database

import (
	"fmt"
	"net"
	"time"

	mysqldriver "github.com/go-sql-driver/mysql"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"ijendev-backend/internal/config"
	"ijendev-backend/internal/models"
)

// Connect membuka koneksi ke MySQL dan memastikan database bisa dihubungi.
func Connect(cfg *config.Config) (*gorm.DB, error) {
	logLevel := logger.Warn
	if !cfg.IsProduction() {
		logLevel = logger.Info // tampilkan query SQL saat development
	}

	db, err := gorm.Open(gormmysql.Open(buildDSN(cfg)), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
		// Agar error "nilai unik bentrok" (misalnya slug kembar) dikenali sebagai gorm.ErrDuplicatedKey.
		TranslateError: true,
	})
	if err != nil {
		return nil, fmt.Errorf("membuka koneksi: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("mengambil koneksi sql: %w", err)
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("ping database %q: %w", cfg.DBName, err)
	}
	return db, nil
}

// Migrate membuat/memperbarui tabel sesuai struct di package models.
// Tabel baru otomatis dibuat; kolom baru otomatis ditambahkan. Kolom lama tidak pernah dihapus otomatis.
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Post{},
	)
}

// buildDSN menyusun connection string MySQL. Memakai mysql.Config agar password
// dengan karakter khusus (@, /, dll) tetap aman.
func buildDSN(cfg *config.Config) string {
	dsn := mysqldriver.NewConfig()
	dsn.User = cfg.DBUser
	dsn.Passwd = cfg.DBPassword
	dsn.Net = "tcp"
	dsn.Addr = net.JoinHostPort(cfg.DBHost, cfg.DBPort)
	dsn.DBName = cfg.DBName
	dsn.Collation = "utf8mb4_unicode_ci"
	dsn.ParseTime = true
	dsn.Loc = time.UTC // waktu disimpan dalam UTC; frontend yang menampilkannya sesuai zona waktu
	return dsn.FormatDSN()
}
