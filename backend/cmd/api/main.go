// Command api adalah titik mulai server backend Ijendev.
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ijendev-backend/internal/config"
	"ijendev-backend/internal/database"
	"ijendev-backend/internal/routes"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg := config.Load()
	if err := cfg.Validate(); err != nil {
		return err
	}
	if err := os.MkdirAll(cfg.UploadDir, 0o755); err != nil {
		return fmt.Errorf("membuat folder upload %q: %w", cfg.UploadDir, err)
	}

	db, err := database.Connect(cfg)
	if err != nil {
		return fmt.Errorf("koneksi database gagal: %w", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("mengambil koneksi sql: %w", err)
	}
	defer sqlDB.Close()

	if cfg.AutoMigrate {
		if err := database.Migrate(db); err != nil {
			return fmt.Errorf("migrasi tabel gagal: %w", err)
		}
		log.Println("migrasi tabel selesai")
	}

	server := &http.Server{
		Addr:              ":" + cfg.AppPort,
		Handler:           routes.Setup(cfg, db),
		ReadHeaderTimeout: 10 * time.Second,
	}

	// Jalankan server di goroutine agar main bisa menunggu sinyal Ctrl+C.
	serverErr := make(chan error, 1)
	go func() {
		log.Printf("server berjalan di http://localhost:%s (env: %s)", cfg.AppPort, cfg.AppEnv)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErr:
		return fmt.Errorf("server berhenti: %w", err)
	case <-quit:
		log.Println("mematikan server...")
	}

	// Beri waktu request yang sedang berjalan untuk selesai.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return server.Shutdown(ctx)
}
