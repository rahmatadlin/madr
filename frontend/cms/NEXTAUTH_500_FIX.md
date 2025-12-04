# Fix NextAuth 500 Error - Session Endpoint

## 🔍 Masalah

Error `GET http://localhost:3000/api/auth/session 500 (Internal Server Error)` terjadi karena NextAuth route handler mengembalikan error.

## ✅ Solusi yang Sudah Diterapkan

1. **Route Handler** - Sudah menggunakan format yang benar untuk NextAuth v5 beta
2. **Session Callback** - Sudah ditambahkan null checks untuk menghindari error
3. **Debug Mode** - Sudah diaktifkan untuk development

## 🔧 Langkah Debugging

### 1. Cek Terminal Logs

Saat dev server running, cek terminal untuk error message spesifik dari NextAuth. Error biasanya muncul saat:
- `NEXTAUTH_SECRET` tidak di-set atau invalid
- Ada error di `authOptions` callbacks
- Ada error di route handler

### 2. Test dengan Verbose Output

```bash
curl -v http://localhost:3000/api/auth/session 2>&1 | grep -A 10 "< HTTP"
```

### 3. Cek Browser Console

Buka browser console dan cek:
- Error message yang lebih spesifik
- Network tab untuk melihat response body dari `/api/auth/session`

### 4. Cek Environment Variables

Pastikan `.env` file berisi:

```env
NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1
NEXTAUTH_SECRET=<generated-secret>
NEXTAUTH_URL=http://localhost:3000
```

**Generate NEXTAUTH_SECRET:**
```bash
openssl rand -base64 32
```

### 5. Restart Dev Server

Setelah mengubah `.env` atau kode, **RESTART** dev server:

```bash
# Stop dev server (Ctrl+C)
cd frontend/cms
npm run dev
```

## 🐛 Common Issues

### Issue 1: Session Callback Error

**Error:** `Cannot read property 'user' of undefined`

**Solution:**
- Pastikan null checks di session callback
- Pastikan `session` dan `token` tidak undefined

### Issue 2: Route Handler Type Error

**Error:** `Type error in route handler`

**Solution:**
- Pastikan menggunakan `as any` untuk compatibility dengan NextAuth v5 beta
- Atau update ke NextAuth v5 stable jika sudah tersedia

### Issue 3: NEXTAUTH_SECRET Missing

**Error:** `[next-auth][error][NO_SECRET]`

**Solution:**
1. Generate secret: `openssl rand -base64 32`
2. Set di `.env`: `NEXTAUTH_SECRET=<generated-secret>`
3. Restart dev server

## 📋 Checklist

- [ ] `.env` file ada dan berisi `NEXTAUTH_SECRET` (sudah di-generate)
- [ ] `NEXTAUTH_URL=http://localhost:3000` sudah benar
- [ ] Dev server sudah di-restart setelah mengubah `.env` atau kode
- [ ] Route handler `/api/auth/[...nextauth]/route.ts` benar
- [ ] Session callback sudah ada null checks
- [ ] Tidak ada error di terminal saat dev server start
- [ ] Backend running di `http://localhost:8080`

## 🆘 Masih Error?

Jika masih error setelah semua langkah di atas:

1. **Cek terminal logs** - Copy error message lengkap dari dev server
2. **Cek browser Network tab** - Lihat response body dari `/api/auth/session`
3. **Test dengan curl verbose** - `curl -v http://localhost:3000/api/auth/session`

Dengan informasi ini, kita bisa debug lebih spesifik.

## 📝 Next Steps

Setelah error teratasi:
1. Test login dari browser
2. Cek session setelah login
3. Test protected routes

