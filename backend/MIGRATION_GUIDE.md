# Database Migration & Seeding Guide

Panduan untuk menjalankan migrasi database dan seeding data awal.

## 📋 Prerequisites

- Database PostgreSQL sudah berjalan
- Environment variables sudah dikonfigurasi di `.env`

## 🚀 Menjalankan Migrasi

### 1. Migrasi ke atas (Up)

Menjalankan semua pending migrations:

```bash
cd backend
go run cmd/migrate/main.go -command=up
```

Atau dari root project:

```bash
go run backend/cmd/migrate/main.go -command=up
```

### 2. Rollback Migrasi (Down)

Rollback satu migration terakhir:

```bash
go run cmd/migrate/main.go -command=down
```

### 3. Cek Versi Migrasi

Cek versi migration saat ini:

```bash
go run cmd/migrate/main.go -command=version
```

## 🌱 Menjalankan Seeding

Menjalankan semua seeders untuk data awal:

```bash
cd backend
go run cmd/seed/main.go
```

Atau dari root project:

```bash
go run backend/cmd/seed/main.go
```

### Data yang di-seed:

1. **Default Admin User**

   - Username: `admin`
   - Password: `admin123`
   - Email: `admin@madr.local`
   - Role: `admin`

2. **Donation Categories**
   - Pembangunan
   - Operasional
   - Sosial
   - Anak Yatim

## 📁 Struktur Migrasi

```
backend/
├── migrations/
│   ├── 000001_create_users_table.up.sql
│   ├── 000001_create_users_table.down.sql
│   ├── 000002_create_refresh_tokens_table.up.sql
│   ├── 000002_create_refresh_tokens_table.down.sql
│   ├── ...
│   └── 000008_create_donations_table.up.sql
│   └── 000008_create_donations_table.down.sql
├── cmd/
│   ├── migrate/
│   │   └── main.go          # CLI untuk migrasi
│   └── seed/
│       └── main.go          # CLI untuk seeding
└── pkg/
    ├── migrate/
    │   └── migrate.go       # Migration utilities
    └── seed/
        └── seed.go          # Seeder functions
```

## 🔄 Workflow Development

### Setup Database Baru

1. **Jalankan migrations:**

   ```bash
   go run cmd/migrate/main.go -command=up
   ```

2. **Jalankan seeding:**

   ```bash
   go run cmd/seed/main.go
   ```

3. **Jalankan server:**
   ```bash
   go run cmd/server/main.go
   ```

### Development Mode

Di development mode (`SERVER_MODE=debug`), server akan otomatis menjalankan AutoMigrate saat startup. Namun, untuk production, selalu gunakan migration files.

## 🐳 Docker

Jika menggunakan Docker, jalankan migrations sebelum start container:

```bash
# Build image
docker build -t madr-backend ./backend

# Run migrations
docker run --rm --env-file backend/.env madr-backend go run cmd/migrate/main.go -command=up

# Run seeding
docker run --rm --env-file backend/.env madr-backend go run cmd/seed/main.go

# Start server
docker run -p 8080:8080 --env-file backend/.env madr-backend
```

## 📝 Membuat Migration Baru

Untuk membuat migration baru secara manual:

1. Buat file `000009_your_migration_name.up.sql` dan `000009_your_migration_name.down.sql` di folder `migrations/`
2. File `.up.sql` berisi SQL untuk apply migration
3. File `.down.sql` berisi SQL untuk rollback migration

## ⚠️ Catatan Penting

- **Jangan hapus migration files** yang sudah dijalankan di production
- **Selalu backup database** sebelum menjalankan migrations di production
- **Test migrations** di staging environment terlebih dahulu
- Seeder akan **skip** jika data sudah ada (idempotent)

## 🔍 Troubleshooting

### Error: "migrations directory not found"

Pastikan Anda menjalankan command dari folder `backend/` atau pastikan path migrations benar.

### Error: "database is in dirty state"

Jika migration gagal di tengah jalan, database akan dalam "dirty state". Untuk memperbaikinya:

1. **Cek versi migration:**

   ```bash
   go run cmd/migrate/main.go -command=version
   ```

2. **Perbaiki masalah di migration file** (jika ada syntax error)

3. **Force version** untuk membersihkan dirty state:

   ```bash
   # Force ke versi sebelumnya (misalnya versi 4 jika gagal di versi 5)
   go run cmd/migrate/main.go -command=force -version=4

   # Atau force ke versi yang sama jika tabel sudah dibuat dengan benar
   go run cmd/migrate/main.go -command=force -version=5
   ```

4. **Jalankan migration lagi:**
   ```bash
   go run cmd/migrate/main.go -command=up
   ```

**Catatan:** Force version hanya membersihkan dirty flag, tidak menjalankan migration. Pastikan struktur database sudah sesuai dengan versi yang di-force.

### Seeder tidak berjalan

Seeder akan skip jika data sudah ada. Untuk re-seed, hapus data terlebih dahulu atau modifikasi seeder untuk force update.
