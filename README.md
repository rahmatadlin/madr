# Masjid Al-Madr - Management System

Monorepo untuk sistem manajemen masjid dengan backend Go, frontend landing page, dan CMS admin dashboard.

## 🏗️ Project Structure

```
madr/
├── backend/              # Go backend API (Gin, GORM, PostgreSQL)
│   ├── cmd/
│   ├── internal/
│   └── pkg/
├── frontend/
│   ├── web/             # Landing page (Next.js 15)
│   └── cms/             # Admin dashboard (Next.js 15)
├── docs/                # Documentation
├── package.json         # Root workspace config (npm workspaces)
└── docker-compose.yml   # Docker setup
```

## 🚀 Quick Start

### Prerequisites

- Node.js 18+
- Go 1.21+
- PostgreSQL 14+
- Docker & Docker Compose (optional)

### Installation

```bash
# Clone repository
git clone <repository-url>
cd madr

# Install all frontend dependencies (monorepo - installs untuk web & cms)
cd frontend
npm install

# Setup backend
cd ../backend
cp .env.example .env
# Edit .env with your database config
go mod download

# Setup frontend environment files
cd ../frontend/web
cp .env.example .env
# Edit .env: NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1

cd ../cms
cp .env.example .env
# Edit .env dengan API URL dan NextAuth config
```

### Development

#### Backend

```bash
cd backend
go run cmd/server/main.go
# Backend runs on http://localhost:8080
```

#### Frontend (Monorepo)

```bash
cd frontend

# Run landing page (port 3000)
npm run dev:web
# atau: cd web && npm run dev

# Run CMS dashboard (port 3001)
npm run dev:cms
# atau: cd cms && npm run dev

# Run both simultaneously
npm run dev:all
```

### Docker (All Services)

```bash
docker-compose up -d
```

## 📦 Monorepo Management

Proyek ini menggunakan **npm workspaces** untuk mengelola multiple frontend applications dalam satu repository. Ini menghindari duplikasi `node_modules` dan memudahkan maintenance.

### 🎯 Keuntungan

- **Shared Dependencies**: Dependencies yang sama (react, next, dll) di-install sekali di root
- **Disk Space Savings**: Menghemat ratusan MB dengan menghindari duplikasi
- **Consistent Versions**: Semua workspace menggunakan versi dependency yang sama
- **Easier Maintenance**: Update dependencies di satu tempat

### Available Scripts

Semua commands dijalankan dari folder `frontend/`:

```bash
cd frontend

# Development
npm run dev:web          # Start web landing page (port 3000)
npm run dev:cms          # Start CMS dashboard (port 3001)
npm run dev:all          # Start both simultaneously

# Build
npm run build:web        # Build web
npm run build:cms        # Build CMS
npm run build:all        # Build all workspaces

# Utilities
npm run clean            # Clean all node_modules & builds
npm install              # Install all workspace dependencies
npm run lint:web         # Lint web workspace
npm run lint:cms         # Lint CMS workspace
```

### Adding Dependencies

```bash
cd frontend

# Add to specific workspace
npm install <package> --workspace=web
# atau singkatnya:
npm install <package> -w web

# Add shared dependency (di root frontend)
npm install -w <package>

# Add dev dependency
npm install -D <package> --workspace=cms
```

### Workspace Structure

```
madr/
├── frontend/
│   ├── package.json         # Workspace config
│   ├── package-lock.json    # Lock file
│   ├── .npmrc              # npm config
│   ├── node_modules/       # Shared dependencies (hoisted)
│   ├── web/                # Landing page
│   │   └── package.json    # Web package config
│   └── cms/                # Admin dashboard
│       └── package.json    # CMS package config
```

**Note**: 
- Semua `node_modules` di-hoist ke `frontend/node_modules`
- Tidak ada `node_modules` di `web/` atau `cms/` (semua shared)
- Shared dependencies seperti `react`, `next`, `typescript` tidak duplikat

Lihat [MONOREPO_SETUP.md](./MONOREPO_SETUP.md) untuk detail lengkap tentang monorepo setup dan best practices.

## 🔐 Default Credentials

### Backend Admin
- Username: `admin`
- Password: `admin123`

**⚠️ Change these in production!**

## 📚 Documentation

- [Backend README](./backend/README.md) - Backend API documentation
- [Frontend Web README](./frontend/web/WEB_README.md) - Landing page docs
- [Frontend CMS README](./frontend/cms/README.md) - CMS dashboard docs
- [API Documentation](./docs/api.md) - Complete API reference
- [Monorepo Setup](./MONOREPO_SETUP.md) - Workspace management guide

## 🛠️ Tech Stack

### Backend
- Go 1.21+
- Gin Framework
- GORM
- PostgreSQL
- JWT Authentication
- Zerolog

### Frontend Web
- Next.js 15
- TypeScript
- Tailwind CSS v4
- TanStack Query
- Framer Motion

### Frontend CMS
- Next.js 15
- TypeScript
- Tailwind CSS v4
- shadcn/ui
- NextAuth.js
- TanStack Query
- Recharts

## 📝 Environment Variables

### Backend (.env)
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=masjid_db
JWT_SECRET=your-secret-key
```

### Frontend Web (.env)
```env
NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1
```

### Frontend CMS (.env)
```env
NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1
NEXTAUTH_SECRET=your-secret-key
NEXTAUTH_URL=http://localhost:3001
```

## 🧪 Testing

```bash
# Backend tests
cd backend
go test ./...

# Frontend tests (when added)
cd frontend
npm run test --workspace=web
npm run test --workspace=cms
```

## 🚢 Deployment

### Backend
```bash
cd backend
docker build -f Dockerfile.backend -t masjid-backend .
docker run -p 8080:8080 masjid-backend
```

### Frontend
```bash
cd frontend

# Build
npm run build:web
npm run build:cms

# Deploy to Vercel/Netlify/etc
# Each workspace can be deployed independently
```

## 🤝 Contributing

1. Create feature branch
2. Make changes
3. Test thoroughly
4. Submit PR

## 📄 License

Copyright © 2024 Masjid Al-Madr. All rights reserved.
