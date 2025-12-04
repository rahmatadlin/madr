# Debug Login Issue - Step by Step

Karena backend sudah berfungsi (curl test berhasil), masalahnya kemungkinan di frontend atau konfigurasi.

## 🔍 Langkah Debugging

### 1. Cek Browser Console

Buka browser Developer Tools (F12) dan cek:

**Tab Console:**

- Cari error merah yang muncul saat login
- Screenshot atau copy error message

**Tab Network:**

- Filter: XHR atau Fetch
- Coba login lagi
- Cari request ke `/api/auth/[...nextauth]` atau `/api/v1/auth/login`
- Klik request tersebut dan cek:
  - **Headers**: Apakah request dikirim dengan benar?
  - **Preview/Response**: Apa response yang diterima?
  - **Status Code**: 200, 401, 500, atau error lain?

### 2. Cek Environment Variables

Pastikan file `frontend/cms/.env` berisi:

```env
NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1
NEXTAUTH_SECRET=your-secret-key-change-in-production
NEXTAUTH_URL=http://localhost:3001
```

**Penting:**

- Restart dev server setelah mengubah `.env`
- `NEXT_PUBLIC_*` variables harus di-restart untuk di-load

### 3. Cek CORS Configuration

Pastikan di `backend/.env`:

```env
CORS_ALLOWED_ORIGINS=http://localhost:3000,http://localhost:3001
```

**Catatan:**

- Port `3001` adalah default untuk CMS Next.js
- Jika CMS berjalan di port lain, tambahkan ke CORS
- Restart backend setelah mengubah CORS

### 4. Test dari Browser Console

Buka browser console dan jalankan:

```javascript
// Test API langsung dari browser
fetch("http://localhost:8080/api/v1/auth/login", {
  method: "POST",
  headers: { "Content-Type": "application/json" },
  body: JSON.stringify({ username: "admin", password: "admin123" }),
})
  .then((r) => r.json())
  .then((data) => {
    console.log("Success:", data);
  })
  .catch((err) => {
    console.error("Error:", err);
  });
```

Jika ini berhasil → masalahnya di NextAuth
Jika ini gagal → masalahnya di CORS atau network

### 5. Cek NextAuth Route Handler

Pastikan file `frontend/cms/app/api/auth/[...nextauth]/route.ts` ada dan benar.

### 6. Cek Logs di Terminal

**Backend logs:**

```bash
cd backend
go run cmd/server/main.go
```

Cari error saat login attempt.

**Frontend logs:**

```bash
cd frontend/cms
npm run dev
```

Cek console output untuk error.

## 🐛 Common Issues & Solutions

### Issue 1: CORS Error di Browser Console

**Error:** `Access to fetch at 'http://localhost:8080/...' from origin 'http://localhost:3001' has been blocked by CORS policy`

**Solution:**

1. Tambahkan port CMS ke `CORS_ALLOWED_ORIGINS` di backend `.env`
2. Restart backend

### Issue 2: NextAuth Error

**Error:** `CredentialsSignin` atau error di NextAuth route

**Solution:**

1. Cek `NEXTAUTH_SECRET` sudah di-set
2. Cek `NEXTAUTH_URL` sesuai dengan URL CMS
3. Generate secret baru: `openssl rand -base64 32`

### Issue 3: Network Error

**Error:** `Network Error` atau `Failed to fetch`

**Solution:**

1. Pastikan backend berjalan di `http://localhost:8080`
2. Test dengan curl untuk memastikan backend accessible
3. Cek firewall atau proxy settings

### Issue 4: 401 Unauthorized

**Error:** `401` status code

**Solution:**

1. Pastikan credentials benar: `admin` / `admin123`
2. Pastikan user sudah di-seed: `go run cmd/seed/main.go`
3. Cek database untuk memastikan user exists dan `is_active = true`

## 📋 Checklist

- [ ] Backend berjalan dan accessible (curl test berhasil)
- [ ] Frontend CMS berjalan tanpa error
- [ ] `.env` file sudah di-set dengan benar
- [ ] CORS_ALLOWED_ORIGINS include port CMS
- [ ] Browser console tidak ada CORS error
- [ ] Network tab menunjukkan request ke backend
- [ ] NextAuth secret sudah di-set
- [ ] User admin sudah di-seed di database

## 🆘 Masih Error?

Jika masih error setelah semua langkah di atas:

1. **Screenshot browser console** (tab Console dan Network)
2. **Copy error message** yang muncul
3. **Cek backend logs** saat login attempt
4. **Cek frontend logs** di terminal

Dengan informasi ini, kita bisa debug lebih spesifik.
