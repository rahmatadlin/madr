# Fix NextAuth 500 Error - `/api/auth/session`

## 🔍 Masalah

Error `GET http://localhost:3000/api/auth/session 500 (Internal Server Error)` terjadi karena NextAuth route handler mengembalikan error.

## ✅ Solusi yang Sudah Diterapkan

1. **Route Handler Fixed** - Route handler sudah diperbaiki untuk NextAuth v5 beta
2. **SessionProvider Wrapper** - SessionProvider dipindah ke wrapper component untuk menghindari SSR issues
3. **Debug Mode** - Debug mode ditambahkan untuk development

## 🔧 Langkah Troubleshooting

### 1. Pastikan Environment Variables

File `frontend/cms/.env` harus berisi:

```env
NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1
NEXTAUTH_SECRET=your-secret-key-change-in-production
NEXTAUTH_URL=http://localhost:3000
```

**PENTING:** 
- Generate secret baru: `openssl rand -base64 32`
- Restart dev server setelah mengubah `.env`

### 2. Cek NextAuth Route Handler

File `app/api/auth/[...nextauth]/route.ts` harus benar:

```typescript
import NextAuth from "next-auth";
import { authOptions } from "@/lib/auth";

const handler = NextAuth(authOptions);

export const GET = handler as any;
export const POST = handler as any;
```

### 3. Test Route Handler Langsung

Test endpoint NextAuth:

```bash
curl http://localhost:3000/api/auth/session
```

Expected response (jika belum login):
```json
{}
```

Jika error, cek terminal untuk error message.

### 4. Cek Browser Console

Buka browser console dan cek:
- Error message yang lebih spesifik
- Network tab untuk melihat response body dari `/api/auth/session`

### 5. Cek Terminal Logs

Saat dev server running, cek terminal untuk error dari NextAuth.

## 🐛 Common Issues

### Issue 1: NEXTAUTH_SECRET tidak di-set

**Error:** `[next-auth][error][NO_SECRET]`

**Solution:**
1. Generate secret: `openssl rand -base64 32`
2. Set di `.env`: `NEXTAUTH_SECRET=<generated-secret>`
3. Restart dev server

### Issue 2: NEXTAUTH_URL salah

**Error:** Session tidak bisa dibuat

**Solution:**
- Set `NEXTAUTH_URL=http://localhost:3000` (sesuai port CMS)
- Restart dev server

### Issue 3: Route handler error

**Error:** 500 Internal Server Error

**Solution:**
- Pastikan `authOptions` tidak ada syntax error
- Cek terminal untuk error message spesifik
- Pastikan semua imports benar

## 📋 Checklist

- [ ] `.env` file ada dan berisi `NEXTAUTH_SECRET` dan `NEXTAUTH_URL`
- [ ] `NEXTAUTH_SECRET` sudah di-generate (bukan default)
- [ ] `NEXTAUTH_URL` sesuai dengan port CMS (3000)
- [ ] Dev server sudah di-restart setelah mengubah `.env`
- [ ] Route handler `/api/auth/[...nextauth]/route.ts` benar
- [ ] Tidak ada error di terminal saat dev server start
- [ ] Browser console tidak ada error lain

## 🆘 Masih Error?

Jika masih error setelah semua langkah di atas:

1. **Cek terminal logs** - Copy error message lengkap
2. **Cek browser Network tab** - Lihat response body dari `/api/auth/session`
3. **Test route handler** - `curl http://localhost:3000/api/auth/session`

Dengan informasi ini, kita bisa debug lebih spesifik.

