# Quick Start - CMS Dev Server

## 🚀 Start Dev Server

```bash
cd frontend/cms
npm run dev
```

Server akan berjalan di `http://localhost:3000`

## ✅ Pre-flight Checklist

Sebelum start dev server, pastikan:

### 1. Environment Variables

File `.env` harus ada di `frontend/cms/` dengan isi:

```env
NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1
NEXTAUTH_SECRET=<generate-dengan-openssl-rand-base64-32>
NEXTAUTH_URL=http://localhost:3000
```

**Generate NEXTAUTH_SECRET:**
```bash
openssl rand -base64 32
```

### 2. Backend Running

Pastikan backend sudah running di `http://localhost:8080`:

```bash
cd backend
go run cmd/server/main.go
```

### 3. Database Ready

Pastikan database sudah di-migrate dan di-seed:

```bash
cd backend
go run cmd/migrate/main.go -command=up
go run cmd/seed/main.go
```

## 🧪 Test Setelah Start

### Test NextAuth Session Endpoint

```bash
curl http://localhost:3000/api/auth/session
```

Expected response (belum login):
```json
{}
```

### Test Backend Login

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin123"}'
```

Expected response:
```json
{
  "message": "Login successful",
  "data": {
    "access_token": "...",
    "refresh_token": "...",
    "user": {...}
  }
}
```

## 🐛 Troubleshooting

### Port 3000 Already in Use

```bash
# Kill process di port 3000
lsof -ti:3000 | xargs kill -9

# Atau gunakan port lain
PORT=3001 npm run dev
```

### NextAuth 500 Error

1. Cek `NEXTAUTH_SECRET` sudah di-set (bukan default)
2. Restart dev server setelah mengubah `.env`
3. Cek terminal untuk error message spesifik

### Backend Connection Error

1. Pastikan backend running di `http://localhost:8080`
2. Cek CORS configuration di backend `.env`:
   ```env
   CORS_ALLOWED_ORIGINS=http://localhost:3000,http://localhost:3001
   ```
3. Restart backend setelah mengubah CORS

## 📋 Complete Startup Sequence

```bash
# Terminal 1: Backend
cd backend
go run cmd/server/main.go

# Terminal 2: CMS Frontend
cd frontend/cms
npm run dev

# Terminal 3: Test
curl http://localhost:3000/api/auth/session
curl http://localhost:8080/api/v1/health
```

