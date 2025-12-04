# Fix CORS Issue untuk CMS Login

## 🔍 Masalah

Error CORS terjadi karena:
1. Backend tidak mengizinkan origin dari CMS
2. CORS configuration tidak di-load dengan benar
3. Backend perlu di-restart setelah mengubah `.env`

## ✅ Solusi

### 1. Pastikan CORS_ALLOWED_ORIGINS di `.env`

Edit `backend/.env` dan pastikan ada:

```env
CORS_ALLOWED_ORIGINS=http://localhost:3000,http://localhost:3001
```

**Catatan:**
- Port `3001` adalah default untuk CMS Next.js
- Jika CMS berjalan di port lain, tambahkan port tersebut
- Pisahkan dengan koma (tanpa spasi)

### 2. Restart Backend

**PENTING:** Setelah mengubah `.env`, **RESTART BACKEND**:

```bash
# Stop backend (Ctrl+C)
# Lalu jalankan lagi:
cd backend
go run cmd/server/main.go
```

### 3. Cek Port CMS

Cek di terminal saat menjalankan CMS:

```bash
cd frontend/cms
npm run dev
```

Lihat output, biasanya:
```
- Local:        http://localhost:3001
```

Pastikan port tersebut sudah ada di `CORS_ALLOWED_ORIGINS`.

### 4. Test CORS dari Browser

Buka CMS di browser (misalnya `http://localhost:3001/login`), lalu buka Console dan jalankan:

```javascript
// Test dengan origin yang benar (dari halaman CMS)
fetch('http://localhost:8080/api/v1/auth/login', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({ username: 'admin', password: 'admin123' })
})
  .then(r => r.json())
  .then(console.log)
  .catch(console.error);
```

**Catatan:** Fetch dari console langsung (bukan dari halaman web) akan memiliki origin `null` dan akan selalu gagal CORS. Test harus dilakukan dari halaman CMS yang sebenarnya.

### 5. Cek Backend Logs

Saat backend start, cek log untuk melihat CORS configuration:

```
{"level":"info","allowed_origins":["http://localhost:3000","http://localhost:3001"],"message":"CORS configuration loaded"}
```

Jika tidak muncul, berarti CORS config tidak di-load dengan benar.

## 🐛 Troubleshooting

### Masalah: CORS masih error setelah restart

1. **Cek format `.env`:**
   ```env
   # BENAR:
   CORS_ALLOWED_ORIGINS=http://localhost:3000,http://localhost:3001
   
   # SALAH (ada spasi):
   CORS_ALLOWED_ORIGINS=http://localhost:3000, http://localhost:3001
   ```

2. **Cek apakah backend membaca `.env`:**
   - Pastikan file `.env` ada di folder `backend/`
   - Pastikan tidak ada typo di nama variable

3. **Cek port CMS:**
   - Buka browser di `http://localhost:3001` (atau port CMS)
   - Cek Network tab untuk melihat origin yang digunakan

### Masalah: Port CMS berbeda

Jika CMS berjalan di port selain 3001:

1. Cek port di terminal saat `npm run dev`
2. Tambahkan ke `CORS_ALLOWED_ORIGINS`:
   ```env
   CORS_ALLOWED_ORIGINS=http://localhost:3000,http://localhost:3001,http://localhost:3002
   ```
3. Restart backend

## ✅ Verifikasi

Setelah fix, login dari CMS seharusnya berhasil. Jika masih error:

1. Cek browser Console untuk error message
2. Cek Network tab untuk melihat request/response
3. Cek backend logs untuk melihat apakah request sampai ke backend

