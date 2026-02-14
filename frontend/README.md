# Frontend Monorepo

Frontend monorepo untuk Masjid Al-Madr: **Vue 3**, **Vite**, **TypeScript**, **Tailwind CSS 4**. Menggunakan npm workspaces.

- **web**: situs publik (port 3000)
- **cms**: dashboard admin (port 3001)

Environment: set `VITE_API_URL` (mis. `http://localhost:8080/api/v1`) di `.env` atau `.env.local` di `web/` dan `cms/` jika backend tidak di `http://localhost:8080/api/v1`.

## 📁 Struktur

```
frontend/
├── package.json         # Workspace config
├── package-lock.json   # Lock file
├── .npmrc              # npm config
├── node_modules/       # Semua dependencies di sini (hoisted)
├── web/                # Landing page
│   └── package.json
└── cms/                # Admin dashboard
    └── package.json
```

**Penting**: Semua `node_modules` di-hoist ke `frontend/node_modules`. Tidak ada `node_modules` di `web/` atau `cms/`.

## 🚀 Quick Start

### Installation

```bash
cd frontend
npm install
```

### Development

```bash
# Run web landing page (port 3000)
npm run dev:web

# Run CMS dashboard (port 3001)
npm run dev:cms

# Run both simultaneously
npm run dev:all
```

### Build

```bash
# Build web
npm run build:web

# Build CMS
npm run build:cms

# Build all
npm run build:all
```

## 📦 Managing Dependencies

```bash
# Add to web
npm install <package> -w web

# Add to CMS
npm install <package> -w cms

# Add shared dependency
npm install -w <package>
```

## 🧹 Cleanup

```bash
npm run clean  # Remove all node_modules and builds
```

Lihat [../MONOREPO_SETUP.md](../MONOREPO_SETUP.md) untuk dokumentasi lengkap.

