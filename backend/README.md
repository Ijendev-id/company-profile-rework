# Backend Ijendev (Go + Gin + MySQL)

## Menjalankan

1. Salin `.env.example` menjadi `.env` (kalau belum), isi `DB_USER` dan `DB_PASSWORD`.
2. Isi `JWT_SECRET` di `.env` (wajib, minimal 32 karakter). Buat dengan:
   ```
   node -e "console.log(require('crypto').randomBytes(48).toString('base64url'))"
   ```
3. Unduh dependensi (di folder `backend`):
   ```
   go get github.com/golang-jwt/jwt/v5 golang.org/x/crypto golang.org/x/term
   go mod tidy
   ```
4. Buat akun admin (sekali saja; password diketik lewat prompt):
   ```
   go run ./cmd/seed -name "Nama Kamu" -email emailkamu@contoh.com
   ```
   Lupa password? Jalankan lagi dengan `-reset`.
5. Jalankan server:
   ```
   go run ./cmd/api
   ```

## Daftar endpoint

Semua respons berformat JSON. Error selalu berbentuk `{"error": "pesan"}`
(validasi: `{"error": "data tidak valid", "fields": {"title": "wajib diisi"}}`).

### Publik
| Method | URL | Fungsi |
|---|---|---|
| GET | `/health` | Cek server dan database |
| POST | `/api/v1/auth/login` | Login: body `{"email","password"}`, hasil berisi `token` |
| GET | `/api/v1/posts` | Daftar artikel terbit. Query: `page`, `per_page` (maks 50), `category` (slug), `q` (cari) |
| GET | `/api/v1/posts/:slug` | Detail artikel terbit |
| GET | `/api/v1/categories` | Daftar kategori |
| GET | `/uploads/<file>` | Gambar hasil upload |

### Admin (header `Authorization: Bearer <token>`)
| Method | URL | Fungsi |
|---|---|---|
| GET | `/api/v1/admin/me` | Data admin yang sedang login |
| GET | `/api/v1/admin/posts` | Semua artikel termasuk draft. Query: `status`, `q`, `page`, `per_page` |
| GET | `/api/v1/admin/posts/:id` | Satu artikel lengkap (untuk form edit) |
| POST | `/api/v1/admin/posts` | Buat artikel |
| PUT | `/api/v1/admin/posts/:id` | Ubah artikel (kirim seluruh isinya) |
| DELETE | `/api/v1/admin/posts/:id` | Hapus artikel (soft delete) |
| POST | `/api/v1/admin/categories` | Buat kategori |
| PUT | `/api/v1/admin/categories/:id` | Ubah kategori |
| DELETE | `/api/v1/admin/categories/:id` | Hapus kategori (artikelnya tetap ada) |
| POST | `/api/v1/admin/uploads` | Upload gambar (multipart, field `file`; JPG/PNG/GIF/WebP) |

Isi body artikel (`POST`/`PUT`):
```json
{
  "title": "Judul artikel",
  "slug": "",                 // kosong = dibuat otomatis dari judul
  "excerpt": "Ringkasan",
  "content": "# Isi dalam Markdown",
  "cover_image": "/uploads/xxxx.jpg",
  "status": "draft",          // draft | published
  "meta_title": "",
  "meta_description": "",
  "category_id": null
}
```

## Tes cepat di PowerShell
```powershell
# 1. Login
$body = @{ email = "emailkamu@contoh.com"; password = "passwordmu" } | ConvertTo-Json
$res = Invoke-RestMethod -Method Post -Uri http://localhost:8080/api/v1/auth/login -ContentType "application/json" -Body $body
$token = $res.token

# 2. Buat artikel
$post = @{ title = "Artikel Pertama"; content = "# Halo`n`nIsi artikel."; status = "published" } | ConvertTo-Json
Invoke-RestMethod -Method Post -Uri http://localhost:8080/api/v1/admin/posts -Headers @{ Authorization = "Bearer $token" } -ContentType "application/json" -Body $post

# 3. Lihat di API publik (atau buka di browser)
Invoke-RestMethod http://localhost:8080/api/v1/posts
```

## Struktur

| Folder | Fungsi | Padanan di Laravel |
|---|---|---|
| `cmd/api` | Server API | `public/index.php` |
| `cmd/seed` | Membuat/mengganti password admin | `php artisan db:seed` |
| `internal/config` | Baca `.env` | `config/*.php` |
| `internal/database` | Koneksi MySQL dan migrasi tabel | `database/migrations` |
| `internal/models` | Struct tabel (users, categories, posts) | `app/Models` |
| `internal/auth` | Token JWT dan hash password | `Auth` / `Hash` |
| `internal/handlers` | Penanganan request | `app/Http/Controllers` |
| `internal/middleware` | CORS, cek login, batas percobaan login | `app/Http/Middleware` |
| `internal/routes` | Daftar URL | `routes/api.php` |
