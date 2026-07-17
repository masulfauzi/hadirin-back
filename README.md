# Hadirin Backend

Backend API untuk aplikasi presensi **Hadirin**. Dikonsumsi oleh tiga klien: frontend web (Vue.js), aplikasi Android, dan aplikasi iOS — semuanya lewat REST API berformat JSON yang sama.

## Tech Stack

| Komponen | Teknologi |
|---|---|
| Bahasa | Go 1.22+ |
| Framework HTTP | [Fiber v2](https://gofiber.io) |
| Database | PostgreSQL |
| ORM | [GORM](https://gorm.io) |
| Migration | [golang-migrate](https://github.com/golang-migrate/migrate) — skema dikelola lewat file SQL, BUKAN `AutoMigrate` |
| Autentikasi | JWT (stateless, header `Authorization: Bearer <token>`), password di-hash dengan bcrypt |

---

## Menjalankan Projek

1. Pastikan PostgreSQL sudah jalan dan database `hadirin` sudah dibuat:
   ```bash
   psql -U postgres -c "CREATE DATABASE hadirin;"
   ```
2. Salin `.env.example` menjadi `.env`, lalu sesuaikan `DB_PASSWORD` dan isi `JWT_SECRET` dengan string acak (jangan pakai nilai contoh):
   ```bash
   cp .env.example .env
   openssl rand -base64 32   # tempel hasilnya ke JWT_SECRET
   ```
3. Jalankan:
   ```bash
   go mod tidy
   go run main.go
   ```
   Migration database jalan otomatis setiap start — tidak perlu perintah migrate terpisah.
4. Cek server hidup:
   ```bash
   curl http://localhost:3000/api/v1/health
   ```

Kalau ada error saat start, lihat bagian [Tips & Troubleshooting](#tips--troubleshooting) di bawah.

---

## Arsitektur: Modular per Fitur

**Aturan utama: satu fitur = satu folder di dalam `modules/`.** Jangan bikin folder `handlers/` atau `models/` global — semua file milik satu fitur (model, query, logika bisnis, handler HTTP, routing) dikumpulkan dalam satu folder modul.

### Struktur Projek

```
hadirin-back/
├── .env                    # konfigurasi lokal (TIDAK di-commit, lihat .gitignore)
├── .env.example            # contoh konfigurasi (di-commit)
├── main.go                 # entry point: load config → migrate → connect db → setup routes → listen
├── config/
│   └── config.go           # baca environment variable, satu-satunya sumber konfigurasi
├── database/
│   ├── database.go         # buka koneksi GORM ke PostgreSQL
│   ├── migrate.go          # embed & jalankan file migration SQL saat start
│   └── migrations/         # file migration SQL bernomor urut (naik terus, tidak pernah diedit ulang)
├── utils/
│   └── response.go         # helper Success()/Error() — satu-satunya cara bikin response JSON
├── middleware/
│   └── auth.go             # middleware JWT, dipakai lintas modul (bukan milik satu modul)
├── routes/
│   └── routes.go           # satu-satunya tempat mendaftarkan routes semua modul
└── modules/
    ├── health/              # cek server & database hidup
    │   ├── handler.go
    │   └── routes.go
    ├── auth/                 # register, login, profil (me)
    │   ├── handler.go
    │   ├── service.go
    │   └── routes.go         # tidak punya model.go/repository.go — data user dipinjam dari modul user
    ├── user/                 # CONTOH ACUAN pola modular lengkap
    │   ├── model.go
    │   ├── repository.go
    │   ├── service.go
    │   ├── handler.go
    │   └── routes.go
    ├── role/                 # roles + user_roles (assign/revoke role ke user)
    ├── menu/                 # menus + role_menu_permissions
    ├── division/             # divisions + division_schedules (jadwal per hari)
    ├── harilibur/            # hari_libur
    ├── karyawan/             # karyawan
    ├── presensi/             # presensi (CRUD dasar, tanpa logika GPS check-in/out)
    └── ijin/                 # ijin + status_ijin & jenis_ijin (read-only)
```

> **Modul `user` adalah contoh acuan.** Kalau bingung bagaimana menulis modul baru, tiru persis pola di `modules/user/`.

### Isi Standar Satu Modul

| File | Tanggung jawab | Boleh berisi |
|---|---|---|
| `model.go` | Struct tabel database (GORM) + tag JSON | Definisi struct saja |
| `repository.go` | Semua query ke database | Query GORM saja, TANPA logika bisnis |
| `service.go` | Logika bisnis (validasi, perhitungan, aturan) | Memanggil repository, TANPA import Fiber/HTTP |
| `handler.go` | Terima request HTTP, kembalikan response JSON | Parsing request + panggil service, TANPA query database langsung |
| `routes.go` | Daftarkan endpoint modul ini | Fungsi `RegisterRoutes` saja |

Modul boleh tidak punya semua file kalau memang tidak butuh (lihat `auth`: tidak punya `model.go`/`repository.go` karena datanya milik modul `user`).

### Alur Request

```
Request → routes.go → handler.go → service.go → repository.go → PostgreSQL
                                                                      │
Response JSON ←────────────────────────────────────────────────────────┘
```

### Aturan yang Tidak Boleh Dilanggar

1. `handler.go` **tidak boleh** memanggil GORM/database langsung — harus lewat `service`.
2. `service.go` **tidak boleh** import Fiber — service tidak boleh tahu-menahu soal HTTP.
3. Satu modul **tidak boleh** mengakses `repository` modul lain — kalau modul A butuh data modul B, panggil `Service` milik modul B (contoh: modul `auth` memanggil `user.Service`, bukan query tabel `users` langsung).
4. Kode yang dipakai lintas modul (misal middleware JWT) ditaruh di `middleware/` di root projek, bukan di dalam salah satu modul.

---

## Format Response JSON

Semua endpoint **wajib** mengembalikan format yang sama, dengan 4 field: `success`, `code`, `message`, `data`.

```json
// Sukses
{ "success": true, "code": 200, "message": "Berhasil mengambil data", "data": [ ... ] }

// Gagal
{ "success": false, "code": 404, "message": "Penjelasan errornya", "data": null }
```

- `code` selalu sama dengan HTTP status code response-nya.
- **Semua handler wajib** memakai `utils.Success(c, code, message, data)` / `utils.Error(c, code, message)` dari [utils/response.go](utils/response.go). Jangan pernah menulis `fiber.Map{...}` response manual di handler — kalau format perlu berubah, cukup ubah di satu tempat itu.

---

## Autentikasi (JWT)

- Stateless, tanpa session/cookie — semua klien (web/Android/iOS) mengirim header `Authorization: Bearer <token>`.
- Password disimpan sebagai hash **bcrypt**, tidak pernah dalam bentuk asli. Field `Password` di `model.go` bertag `json:"-"` supaya tidak pernah ikut ter-serialize ke response.
- Endpoint publik (tanpa token): `POST /api/v1/auth/register`, `POST /api/v1/auth/login`.
- Endpoint terlindungi (butuh token): dibungkus `middleware.Protected(cfg.JWTSecret)` di `routes.go` masing-masing modul, contoh lihat [modules/user/routes.go](modules/user/routes.go).
- Di handler yang terlindungi, ID user yang sedang login diambil dengan:
  ```go
  userID := c.Locals("user_id").(uuid.UUID)
  ```
- Semua primary key & foreign key di database bertipe **UUID** (`gen_random_uuid()`), bukan auto-increment. Di kode Go dipakai `uuid.UUID` dari `github.com/google/uuid`.

---

## Database & Migration

- Skema database dikelola lewat file SQL bernomor urut di `database/migrations/`, **bukan** `AutoMigrate` GORM.
- Setiap perubahan skema = satu pasang file: `NNNNNN_deskripsi.up.sql` (terapkan) dan `NNNNNN_deskripsi.down.sql` (batalkan/rollback).
- File migration di-embed ke dalam binary (`//go:embed migrations/*.sql` di [database/migrate.go](database/migrate.go)) dan dijalankan otomatis setiap `main.go` start. golang-migrate mencatat versi terakhir di tabel `schema_migrations`, jadi migration yang sama tidak dijalankan dua kali.
- **JANGAN PERNAH mengedit file migration yang sudah pernah dijalankan** (apalagi yang sudah di-push/di-deploy). Kalau perlu mengubah tabel yang sudah ada, buat file migration baru dengan nomor urut berikutnya (mis. `ALTER TABLE` di `000003_...`).

---

## Cara Menambah Fitur/Modul Baru

Ikuti pola modul `user` sebagai acuan. Contoh menambah fitur **presensi**:

1. Buat folder `modules/presensi/` berisi `model.go`, `repository.go`, `service.go`, `handler.go`, `routes.go` — semua file memakai `package presensi`.
2. Definisikan struct di `model.go` dengan tag `gorm` dan `json` **snake_case**.
3. Buat migration baru dengan nomor urut berikutnya, contoh:
   - `database/migrations/000002_create_presensi_table.up.sql`
   - `database/migrations/000002_create_presensi_table.down.sql`

   Kalau CLI `migrate` terinstal:
   ```bash
   migrate create -ext sql -dir database/migrations -seq create_presensi_table
   ```
4. Lindungi endpoint dengan JWT kalau perlu login, di `routes.go` modul:
   ```go
   presensiGroup := router.Group("/presensi", middleware.Protected(cfg.JWTSecret))
   ```
5. Daftarkan modul baru di [routes/routes.go](routes/routes.go) — cukup tambah satu baris:
   ```go
   presensi.RegisterRoutes(api, db, cfg)
   ```
6. Selesai. `main.go`, `config/`, dan modul lain **tidak perlu diubah sama sekali**.

Kalau modul baru butuh data dari modul lain (misal `rekap` butuh data `presensi`), panggil `Service` milik modul itu (di-inject lewat `RegisterRoutes`) — **jangan** query tabel modul lain langsung dari repository sendiri.

---

## Environment Variables

Lihat [.env.example](.env.example) untuk daftar lengkap. Yang penting:

| Variable | Fungsi |
|---|---|
| `APP_PORT` | Port server (default `3000`) |
| `DB_*` | Kredensial koneksi PostgreSQL |
| `CORS_ORIGINS` | Origin frontend yang diizinkan (comma-separated) — wajib diisi kalau ada domain frontend baru |
| `JWT_SECRET` | String acak minimal 32 karakter, **wajib diisi** — aplikasi menolak start kalau kosong |
| `JWT_EXPIRE_HOURS` | Umur token JWT dalam jam |

`.env` **tidak boleh ikut ter-commit** (sudah ada di `.gitignore`). Kalau mengubah/menambah variable baru, update juga `.env.example` supaya orang lain tahu variable apa saja yang dibutuhkan.

---

## Tips & Troubleshooting

- `connection refused` saat start → PostgreSQL belum jalan, atau `DB_HOST`/`DB_PORT` di `.env` salah.
- `password authentication failed` → `DB_USER`/`DB_PASSWORD` di `.env` tidak cocok dengan PostgreSQL lokal.
- `database "hadirin" does not exist` → database belum dibuat, jalankan `CREATE DATABASE hadirin;`.
- Import `hadirin-back/...` merah di editor → jalankan `go mod tidy` lalu restart language server.
- Aplikasi berhenti dengan `JWT_SECRET wajib diisi di .env` → isi `JWT_SECRET`.
- Endpoint terlindungi selalu 401 padahal token benar → cek header persis `Authorization: Bearer <token>` (satu spasi setelah `Bearer`, token tidak terpotong).
- Token selalu dianggap kedaluwarsa → cek `JWT_EXPIRE_HOURS` bukan 0/negatif.
- `Dirty database version X` → sebuah migration gagal di tengah jalan. Perbaiki penyebabnya dulu, lalu reset penanda:
  ```bash
  migrate -database "postgres://postgres:PASSWORD@localhost:5432/hadirin?sslmode=disable" -path database/migrations force X
  ```
- Rollback manual satu langkah (development saja):
  ```bash
  migrate -database "postgres://postgres:PASSWORD@localhost:5432/hadirin?sslmode=disable" -path database/migrations down 1
  ```
- Bingung menaruh kode di mana? Tanya "kode ini milik fitur apa?" → taruh di `modules/<nama-fitur>/`. Kode yang dipakai semua fitur (config, koneksi database) tetap di `config/` dan `database/`.

---

## Belum Termasuk Scope Saat Ini

Hal-hal berikut sengaja belum dikerjakan, akan menyusul di issue terpisah:

- Modul rekap (ikuti panduan "Cara Menambah Fitur/Modul Baru" di atas saat waktunya tiba)
- Logika bisnis presensi (check-in/check-out dengan validasi GPS, penentuan status hadir dari jadwal divisi)
- Alur approval ijin (perubahan status oleh approver, notifikasi)
- Refresh token, logout, lupa/reset password
- Pembatasan endpoint berdasarkan role & permission menu (RBAC enforcement) — tabelnya sudah ada, enforcement menyusul
- Unit test
- Dockerfile / deployment
- Dokumentasi API (Swagger)

Untuk detail rencana awal setup projek ini, lihat `issue.md` (tidak ikut ter-commit, hanya arsip lokal).
