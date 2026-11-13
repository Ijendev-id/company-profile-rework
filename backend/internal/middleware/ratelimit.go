package middleware

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// RateLimit membatasi jumlah request per alamat IP dalam rentang waktu tertentu
// (dipakai untuk endpoint login agar password tidak bisa ditebak berulang-ulang).
// Penghitungan disimpan di memori, jadi ikut reset saat server dimulai ulang.
//
// Catatan: kalau login dipanggil lewat server Nuxt, semua request tampak berasal dari IP server Nuxt,
// sehingga batasnya berlaku bersama untuk semua pengguna. Untuk satu admin ini justru cukup aman.
func RateLimit(max int, window time.Duration) gin.HandlerFunc {
	var mu sync.Mutex
	hits := make(map[string][]time.Time)

	return func(c *gin.Context) {
		ip := c.ClientIP()
		now := time.Now()
		cutoff := now.Add(-window)

		mu.Lock()
		recent := hits[ip][:0]
		for _, t := range hits[ip] {
			if t.After(cutoff) {
				recent = append(recent, t)
			}
		}

		if len(recent) >= max {
			hits[ip] = recent
			mu.Unlock()
			c.Header("Retry-After", strconv.Itoa(int(window.Seconds())))
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "terlalu banyak percobaan, coba lagi nanti"})
			return
		}

		hits[ip] = append(recent, now)

		// Bersihkan IP lama supaya map tidak membesar tanpa batas.
		if len(hits) > 10000 {
			for key, times := range hits {
				if len(times) == 0 || times[len(times)-1].Before(cutoff) {
					delete(hits, key)
				}
			}
		}
		mu.Unlock()

		c.Next()
	}
}
