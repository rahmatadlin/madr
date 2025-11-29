# Masjid Management System

Sistem manajemen masjid lengkap dengan landing page publik, CMS admin dashboard, dan backend REST API.

## 🎯 Project Overview

Masjid Management System adalah aplikasi full-stack untuk mengelola berbagai aspek masjid, termasuk:

- Pengumuman dan berita
- Event dan kegiatan
- Galeri foto
- Donasi
- Banner dan konten landing page

## 🏗️ Tech Stack

### Frontend

- **Next.js 15** (App Router)
- **TypeScript**
- **Tailwind CSS v4**
- **shadcn/ui**
- **TanStack Query**
- **React Hook Form + Zod**
- **Framer Motion**

### Backend

- **Golang 1.21+**
- **Gin Framework**
- **GORM ORM**
- **PostgreSQL**
- **JWT Authentication**
- **Zerolog** (Structured Logging)

### DevOps

- **Docker & Docker Compose**
- **GitHub Actions** (CI/CD - coming soon)

## 📦 Monorepo Structure

```
madr/
├── backend/          # Golang REST API
│   ├── cmd/         # Application entry points
│   ├── internal/    # Internal packages
│   │   ├── config/  # Configuration
│   │   ├── domain/  # Domain models
│   │   ├── repository/ # Data access layer
│   │   ├── usecase/ # Business logic
│   │   ├── handler/ # HTTP handlers
│   │   └── middleware/ # HTTP middlewares
│   ├── pkg/         # Shared packages
│   └── migrations/  # Database migrations
├── frontend/        # Frontend applications
│   ├── web/         # Public landing page (Next.js)
│   └── cms/         # Admin dashboard (Next.js)
├── docs/            # Documentation
│   └── api.md       # API documentation
└── docker-compose.yml # Docker compose configuration
```

## 🚀 Quick Start

### Prerequisites

- Docker & Docker Compose
- Go 1.21+ (untuk development lokal)
- Node.js 18+ (untuk frontend - coming soon)

### Setup dengan Docker Compose

1. **Clone repository:**

```bash
git clone <repository-url>
cd madr
```

2. **Setup environment variables:**

```bash
cd backend
cp .env.example .env
# Edit .env sesuai kebutuhan
cd ..
```

3. **Jalankan semua services:**

```bash
docker-compose up -d
```

4. **Cek status:**

```bash
docker-compose ps
```

5. **Cek logs:**

```bash
docker-compose logs -f backend
```

6. **Test API:**

```bash
curl http://localhost:8080/health
```

### Setup Development Lokal

#### Backend

1. **Masuk ke folder backend:**

```bash
cd backend
```

2. **Install dependencies:**

```bash
go mod download
```

3. **Setup database:**

```bash
# Menggunakan Docker
docker-compose up -d postgres
```

4. **Setup environment:**

```bash
cp .env.example .env
# Edit .env
```

5. **Jalankan server:**

```bash
go run cmd/server/main.go
```

Lihat [backend/README.md](./backend/README.md) untuk detail lebih lanjut.

#### Frontend

> **Coming soon** - Frontend akan diimplementasikan pada tahap selanjutnya.

## 📚 Documentation

- [Backend README](./backend/README.md) - Dokumentasi backend API
- [API Documentation](./docs/api.md) - Dokumentasi lengkap API endpoints

## 🔌 API Endpoints

### Base URL

```
http://localhost:8080/api/v1
```

### Available Endpoints

#### Health Check

- `GET /health` - Health check

#### Announcements

- `GET /announcements/published` - Get published announcements (Public)
- `GET /announcements/:id` - Get announcement by ID (Public)
- `POST /admin/announcements` - Create announcement (Admin)
- `GET /admin/announcements` - Get all announcements (Admin)
- `PUT /admin/announcements/:id` - Update announcement (Admin)
- `DELETE /admin/announcements/:id` - Delete announcement (Admin)

Lihat [docs/api.md](./docs/api.md) untuk dokumentasi lengkap.

## 🧪 Testing

### Test Health Check

```bash
curl http://localhost:8080/health
```

### Test Create Announcement

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

### Test Get Published Announcements

```bash
curl http://localhost:8080/api/v1/announcements/published
```

## 🏛️ Architecture

### Backend Architecture

Backend menggunakan **Clean Architecture** dengan pemisahan layer:

1. **Domain Layer**: Business entities dan models
2. **Repository Layer**: Data access abstraction
3. **UseCase Layer**: Business logic
4. **Handler Layer**: HTTP request handling

### Frontend Architecture

> **Coming soon**

## 🔒 Security Features

- Rate limiting per IP
- CORS configuration
- Error handling middleware
- Structured logging
- JWT authentication (coming soon)

## 🐳 Docker

### Build Backend Image

```bash
docker build -t madr-backend ./backend
```

### Run dengan Docker Compose

```bash
docker-compose up -d
```

### Stop Services

```bash
docker-compose down
```

### Stop dan Hapus Volumes

```bash
docker-compose down -v
```

## 📝 Environment Variables

### Backend

Lihat `backend/.env.example` untuk daftar lengkap environment variables.

**Important variables:**

- `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`
- `JWT_SECRET` (ubah di production!)
- `CORS_ALLOWED_ORIGINS`
- `RATE_LIMIT_*`

## 🗄️ Database

### PostgreSQL

Default configuration:

- Host: `localhost`
- Port: `5432`
- User: `postgres`
- Password: `postgres`
- Database: `madr_db`

### Migrations

Saat ini menggunakan GORM AutoMigrate untuk development. Untuk production, akan digunakan migration tool seperti golang-migrate.

## 📋 Development Roadmap

### ✅ Completed

- [x] Backend structure dengan Clean Architecture
- [x] Gin REST API setup
- [x] Database connection (PostgreSQL)
- [x] CRUD module Announcement
- [x] Middleware (error handling, rate limiting, CORS)
- [x] Structured logging
- [x] Docker setup
- [x] Documentation

### 🚧 In Progress

- [ ] JWT Authentication
- [ ] Frontend Web (Landing Page)
- [ ] Frontend CMS (Admin Dashboard)

### 📅 Planned

- [ ] CRUD Events
- [ ] CRUD Donations
- [ ] CRUD Gallery
- [ ] CRUD Banner
- [ ] File upload handling
- [ ] Unit tests & Integration tests
- [ ] CI/CD pipeline
- [ ] API documentation dengan Swagger

## 🤝 Contributing

1. Fork repository
2. Create feature branch (`git checkout -b feature/amazing-feature`)
3. Commit changes (`git commit -m 'feat: add amazing feature'`)
4. Push to branch (`git push origin feature/amazing-feature`)
5. Open Pull Request

### Commit Message Convention

Menggunakan [Conventional Commits](https://www.conventionalcommits.org/):

- `feat:` - New feature
- `fix:` - Bug fix
- `docs:` - Documentation changes
- `style:` - Code style changes (formatting, etc.)
- `refactor:` - Code refactoring
- `test:` - Adding tests
- `chore:` - Maintenance tasks

**Examples:**

```
feat: add JWT authentication middleware
fix: resolve database connection timeout
docs: update API documentation
refactor: improve error handling in usecase layer
```

## 📄 License

MIT License

## 👥 Authors

- Development Team

## 🙏 Acknowledgments

- Gin Framework community
- GORM community
- Next.js team

---

**Status**: 🚧 Development in Progress
