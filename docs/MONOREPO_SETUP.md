# Monorepo Setup Guide

Proyek ini menggunakan **npm workspaces** untuk mengelola multiple frontend applications dalam satu repository.

## 📁 Struktur Monorepo

```
madr/
├── frontend/
│   ├── package.json         # Workspace config
│   ├── package-lock.json    # Lock file
│   ├── .npmrc              # npm workspace config
│   ├── node_modules/       # Shared dependencies (hoisted)
│   ├── web/                # Landing page (Next.js)
│   │   ├── package.json
│   │   └── ...             # NO node_modules (semua di root)
│   └── cms/                # Admin dashboard (Next.js)
│       ├── package.json
│       └── ...             # NO node_modules (semua di root)
└── backend/                # Go backend (terpisah)
```

## 🚀 Keuntungan Monorepo

### 1. **Shared Dependencies**

Dependencies yang sama (seperti `react`, `next`, `typescript`) di-install sekali di root `node_modules`, bukan duplikat di setiap workspace.

**Sebelum (tanpa workspaces):**

```
frontend/web/node_modules/    → 684MB
frontend/cms/node_modules/    → 728MB
Total: ~1.4GB
```

**Sesudah (dengan workspaces):**

```
node_modules/ (root)          → ~14MB (shared deps)
frontend/web/node_modules/    → Hanya deps spesifik web
frontend/cms/node_modules/    → Hanya deps spesifik cms
Total: Lebih kecil & efisien!
```

### 2. **Consistent Versions**

Semua workspace menggunakan versi dependency yang sama, menghindari version conflicts.

### 3. **Easier Maintenance**

Update dependencies di satu tempat, semua workspace otomatis ter-update.

### 4. **Faster CI/CD**

Build dan test bisa dijalankan parallel untuk multiple workspaces.

### 5. **Code Sharing** (Future)

Bisa share utilities/types antar packages dengan membuat shared package.

## 📦 Setup & Installation

### Initial Setup

```bash
# Masuk ke folder frontend
cd frontend

# Install semua dependencies (untuk web & cms)
npm install

# npm akan otomatis:
# 1. Install semua deps di frontend/node_modules (hoisted)
# 2. Tidak ada node_modules di web/ atau cms/ (semua shared)
# 3. Create symlinks jika diperlukan
```

### Development

```bash
# Masuk ke folder frontend
cd frontend

# Run web landing page
npm run dev:web
# atau
cd web && npm run dev

# Run CMS dashboard
npm run dev:cms
# atau
cd cms && npm run dev

# Run both simultaneously (menggunakan concurrently)
npm run dev:all
```

### Build

```bash
cd frontend

# Build web
npm run build:web

# Build CMS
npm run build:cms

# Build all
npm run build:all
```

## 🔧 How Workspaces Work

### Dependency Hoisting

npm workspaces secara otomatis "hoist" semua dependencies ke `frontend/node_modules`:

```
frontend/node_modules/
├── react/          # Shared by web & cms (hoisted)
├── react-dom/      # Shared (hoisted)
├── next/           # Shared (hoisted)
├── typescript/     # Shared (hoisted)
└── ...             # Semua dependencies di sini

frontend/web/
└── (NO node_modules - semua di root)

frontend/cms/
└── (NO node_modules - semua di root)
```

**Semua dependencies di-hoist ke `frontend/node_modules`**, tidak ada duplikasi!

### Adding Dependencies

```bash
cd frontend

# Add to specific workspace
npm install axios --workspace=web
# atau singkatnya:
npm install axios -w web

# Add shared dependency (di root frontend)
npm install -w typescript

# Add dev dependency
npm install -D eslint --workspace=cms
```

### Removing Dependencies

```bash
cd frontend
npm uninstall axios --workspace=web
```

### Checking Dependencies

```bash
cd frontend

# List all dependencies across workspaces
npm list --workspaces

# Check specific package version
npm list react --workspaces

# Check dependencies of specific workspace
npm list --workspace=web
```

## 📝 Best Practices

### 1. **Shared Dependencies**

Letakkan di root jika digunakan oleh multiple workspaces:

- `react`, `react-dom` (harus sama versi)
- `typescript`, `eslint` (shared tooling)
- `axios`, `zod` (shared utilities)

### 2. **Package-Specific**

Letakkan di workspace masing-masing:

- `next` (bisa beda versi jika perlu)
- Package-specific configs
- Unique dependencies (misal: `next-auth` hanya di CMS)

### 3. **Version Consistency**

Gunakan versi yang sama untuk shared deps:

```json
// Di root package.json (opsional, untuk enforce)
"overrides": {
  "react": "^19.2.0",
  "react-dom": "^19.2.0"
}
```

### 4. **Scripts**

Gunakan root scripts untuk convenience:

- `npm run dev:web` lebih mudah daripada `cd frontend/web && npm run dev`
- `npm run build:all` untuk build semua sekaligus

## 🔄 Migration dari Separate Projects

Jika sudah ada project terpisah (seperti sekarang):

1. ✅ **Keep existing structure**: Tidak perlu pindah file
2. ✅ **Add root package.json**: Dengan workspaces config
3. ✅ **Run npm install**: Di root untuk hoist dependencies
4. ✅ **Update scripts**: Gunakan workspace commands

**Tidak ada breaking changes!** Setiap workspace tetap bisa dijalankan secara independen.

## 🐛 Troubleshooting

### Dependencies tidak ter-hoist

```bash
cd frontend

# Clear semua node_modules
npm run clean

# Reinstall
npm install
```

### Version conflicts

```bash
cd frontend

# Check installed versions
npm list react --workspaces

# Force resolution (di frontend/package.json)
"overrides": {
  "react": "^19.2.0",
  "react-dom": "^19.2.0"
}
```

### Build errors

```bash
cd frontend

# Clear Next.js cache
npm run clean

# Rebuild
npm run build:all
```

### "Module not found" errors

Pastikan sudah run `npm install` di root setelah menambah dependency baru.

## 📊 Space Savings

Dengan workspaces, kita menghemat:

- **Disk space**: Shared deps tidak duplikat
- **Install time**: Install sekali untuk shared deps
- **CI/CD time**: Cache bisa di-share

## 🔮 Future Enhancements

### Shared Packages

Bisa membuat shared package untuk code yang digunakan kedua workspace:

```
frontend/
├── web/
├── cms/
└── shared/              # Shared utilities
    ├── package.json
    ├── types/
    └── utils/
```

Kemudian import di workspace:

```typescript
import { sharedUtil } from "@madr/shared";
```

### pnpm Workspaces (Alternative)

Jika ingin lebih efisien lagi, bisa migrasi ke pnpm:

1. Install pnpm: `npm install -g pnpm`
2. Update `package.json` dengan pnpm config
3. Run `pnpm install`

**pnpm advantages:**

- Hard links untuk disk efficiency
- Faster installs
- Better dependency resolution
- Content-addressable storage

## 🎯 Current Status

✅ **Setup Complete**

- npm workspaces configured di `frontend/`
- `frontend/package.json` dengan scripts
- Semua dependencies di-hoist ke `frontend/node_modules`
- **Tidak ada `node_modules` di `web/` atau `cms/`** (semua shared)

✅ **Working Commands** (semua dari folder `frontend/`)

- `npm run dev:web` - Start web (port 3000)
- `npm run dev:cms` - Start CMS (port 3001)
- `npm run dev:all` - Start both simultaneously
- `npm run build:all` - Build all workspaces
- `npm run clean` - Clean all node_modules & builds

✅ **Benefits Achieved**

- Reduced disk usage (semua deps di satu tempat: `frontend/node_modules`)
- Consistent dependencies (satu versi untuk semua)
- Easier maintenance (update di satu tempat)
- Centralized scripts di `frontend/`

## ✅ Struktur Final

```
madr/
├── frontend/
│   ├── package.json         ✅ Workspace config
│   ├── package-lock.json    ✅ Lock file
│   ├── .npmrc              ✅ npm config
│   ├── node_modules/       ✅ Semua dependencies (586MB)
│   ├── web/                ✅ Landing page
│   │   └── package.json   ✅ (NO node_modules)
│   └── cms/                ✅ Admin dashboard
│       └── package.json   ✅ (NO node_modules)
└── backend/                # Go backend (terpisah)
```

**Cara Pakai:**

```bash
cd frontend
npm install          # Install semua dependencies (sekali)
npm run dev:all      # Run web & CMS bersamaan
npm run build:all    # Build semua workspace
```

## 📚 References

- [npm workspaces docs](https://docs.npmjs.com/cli/v9/using-npm/workspaces)
- [npm workspaces tutorial](https://docs.npmjs.com/cli/v9/using-npm/workspaces#workspaces-tutorial)
