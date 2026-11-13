// Package auth berisi pembuatan/pengecekan token login (JWT) dan password (bcrypt).
package auth

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const issuer = "ijendev"

// TokenManager membuat dan memeriksa token JWT yang ditandatangani dengan JWT_SECRET.
type TokenManager struct {
	secret []byte
	ttl    time.Duration
}

func NewTokenManager(secret string, ttl time.Duration) *TokenManager {
	return &TokenManager{secret: []byte(secret), ttl: ttl}
}

// Generate membuat token untuk user dan mengembalikan waktu kedaluwarsanya.
func (m *TokenManager) Generate(userID uint) (string, time.Time, error) {
	now := time.Now()
	expiresAt := now.Add(m.ttl)

	claims := jwt.RegisteredClaims{
		Issuer:    issuer,
		Subject:   strconv.FormatUint(uint64(userID), 10),
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(expiresAt),
	}

	signed, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(m.secret)
	if err != nil {
		return "", time.Time{}, err
	}
	return signed, expiresAt, nil
}

// Parse memeriksa tanda tangan, penerbit, dan masa berlaku token, lalu mengembalikan ID user.
// Hanya algoritma HS256 yang diterima (mencegah serangan pergantian algoritma).
func (m *TokenManager) Parse(tokenString string) (uint, error) {
	claims := &jwt.RegisteredClaims{}
	_, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(*jwt.Token) (any, error) { return m.secret, nil },
		jwt.WithValidMethods([]string{"HS256"}),
		jwt.WithIssuer(issuer),
		jwt.WithExpirationRequired(),
	)
	if err != nil {
		return 0, err
	}

	id, err := strconv.ParseUint(claims.Subject, 10, 32)
	if err != nil || id == 0 {
		return 0, errors.New("subject token tidak valid")
	}
	return uint(id), nil
}
