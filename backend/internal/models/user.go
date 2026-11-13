// Package models berisi struct yang dipetakan menjadi tabel database (mirip Model + migration di Laravel).
package models

import "time"

// User adalah akun admin. Saat ini hanya ada satu admin.
type User struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"size:100;not null" json:"name"`
	Email string `gorm:"size:191;uniqueIndex;not null" json:"email"`

	// Hash bcrypt, bukan password asli. json:"-" memastikan tidak pernah ikut terkirim di response API.
	PasswordHash string `gorm:"size:255;not null" json:"-"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
