package auth

import "golang.org/x/crypto/bcrypt"

// MaxPasswordBytes adalah batas panjang password yang didukung bcrypt.
const MaxPasswordBytes = 72

// HashPassword mengubah password menjadi hash bcrypt untuk disimpan di database.
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CheckPassword true kalau password cocok dengan hash.
func CheckPassword(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

var dummyHash []byte

func init() {
	dummyHash, _ = bcrypt.GenerateFromPassword([]byte("dummy-password"), bcrypt.DefaultCost)
}

// FakeCheck menjalankan perbandingan bcrypt palsu. Dipanggil saat email tidak ditemukan
// supaya waktu respons sama dengan email yang ada (mencegah orang menebak email admin dari kecepatan respons).
func FakeCheck(password string) {
	_ = bcrypt.CompareHashAndPassword(dummyHash, []byte(password))
}
