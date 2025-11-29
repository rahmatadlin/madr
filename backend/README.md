# Backend API - Masjid Management System

Backend REST API untuk Masjid Management System menggunakan Golang, Gin Framework, dan PostgreSQL.

## 🏗️ Arsitektur

Backend menggunakan **Clean Architecture** dengan struktur sebagai berikut:

```
backend/
├── cmd/
│   └── server/          # Entry point aplikasi
├── internal/
│   ├── config/          # Konfigurasi aplikasi
│   ├── domain/          # Domain models & entities
│   │   ├── models/      # Base models
│   │   └── announcement/ # Domain entity
│   ├── repository/      # Data access layer
│   │   └── announcement/
│   ├── usecase/         # Business logic layer
│   │   └── announcement/
│   ├── handler/         # HTTP handlers
│   │   └── announcement/
│   └── middleware/      # HTTP middlewares
│       ├── error_handler.go
│       └── rate_limiter.go
├── pkg/                 # Shared packages
│   ├── database/        # Database connection
│   └── logger/          # Structured logging
└── migrations/          # Database migrations (future)
```

## 🚀 Quick Start

### Prerequisites

- Go 1.21 atau lebih baru
- PostgreSQL 12+ atau MySQL 8+
- Docker & Docker Compose (opsional)

### Setup Lokal

1. **Clone repository dan masuk ke folder backend:**

```bash
cd backend
```

2. **Install dependencies:**

```bash
go mod download
```

3. **Setup environment variables:**

```bash
cp .env.example .env
# Edit .env sesuai kebutuhan
```

4. **Jalankan database (menggunakan Docker):**

```bash
docker-compose up -d postgres
```

5. **Jalankan aplikasi:**

```bash
go run cmd/server/main.go
```

Server akan berjalan di `http://localhost:8080`

### Setup dengan Docker Compose

1. **Jalankan semua services:**

```bash
docker-compose up -d
```

2. **Cek logs:**

```bash
docker-compose logs -f backend
```

3. **Stop services:**

```bash
docker-compose down
```

## 📝 Environment Variables

Lihat `.env.example` untuk daftar lengkap environment variables. Yang penting:

- `SERVER_HOST` & `SERVER_PORT`: Alamat server
- `DB_*`: Konfigurasi database
- `JWT_SECRET`: Secret key untuk JWT (ubah di production!)
- `CORS_ALLOWED_ORIGINS`: Origins yang diizinkan untuk CORS
- `RATE_LIMIT_*`: Konfigurasi rate limiting

## 🔌 API Endpoints

### Health Check

- `GET /health` - Health check endpoint
- `GET /api/health` - Health check dengan detail

### Announcements (Public)

- `GET /api/v1/announcements/published` - Get published announcements
- `GET /api/v1/announcements/:id` - Get announcement by ID

### Announcements (Admin)

- `POST /api/v1/admin/announcements` - Create announcement
- `GET /api/v1/admin/announcements` - Get all announcements
- `PUT /api/v1/admin/announcements/:id` - Update announcement
- `DELETE /api/v1/admin/announcements/:id` - Delete announcement

## 🧪 Testing API

### Health Check

```bash
curl http://localhost:8080/health
```

### Create Announcement

```bash
curl -X POST http://localhost:8080/api/v1/admin/announcements \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Pengumuman Sholat Jumat",
    "content": "Sholat Jumat akan dilaksanakan pada pukul 12:00 WIB",
    "is_published": true,
    "author": "Admin Masjid"
  }'
```

### Get Published Announcements

```bash
curl http://localhost:8080/api/v1/announcements/published?limit=10&offset=0
```

## 🏛️ Clean Architecture Layers

### Domain Layer

Berisi business entities dan models. Tidak bergantung pada framework atau library eksternal.

### Repository Layer

Abstraksi untuk data access. Menggunakan interface untuk memudahkan testing dan perubahan database.

### UseCase Layer

Berisi business logic. Menggunakan repository interface untuk mengakses data.

### Handler Layer

HTTP handlers yang menerima request dan memanggil use case. Menggunakan Gin framework.

## 🔒 Security Features

- **Rate Limiting**: Membatasi jumlah request per IP
- **CORS**: Cross-Origin Resource Sharing configuration
- **Error Handling**: Centralized error handling middleware
- **Structured Logging**: Logging menggunakan zerolog

## 📊 Logging

Backend menggunakan structured logging dengan zerolog. Format logging dapat diatur melalui environment variable:

- `LOG_LEVEL`: debug, info, warn, error
- `LOG_FORMAT`: json (production) atau console (development)

## 🐳 Docker

### Build Image

```bash
docker build -t madr-backend ./backend
```

### Run Container

```bash
docker run -p 8080:8080 --env-file .env madr-backend
```

## 🔄 Database Migrations

Saat ini menggunakan GORM AutoMigrate untuk development. Untuk production, disarankan menggunakan migration tool seperti:

- [golang-migrate](https://github.com/golang-migrate/migrate)
- [gormigrate](https://github.com/go-gormigrate/gormigrate)

## 📚 Next Steps

1. Implementasi JWT authentication
2. Setup proper database migrations
3. Unit tests & integration tests
4. API documentation dengan Swagger/OpenAPI
5. CI/CD pipeline

## 📄 License

MIT
