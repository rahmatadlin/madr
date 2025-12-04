# Troubleshooting CMS Login Issues

## 🔍 Common Login Errors

### 1. Internal Server Error / 500 Error

**Kemungkinan Penyebab:**

- Backend tidak berjalan
- CORS configuration salah
- API URL tidak benar
- Database connection error
- JWT secret tidak di-set

**Solusi:**

1. **Cek Backend Status:**

   ```bash
   curl http://localhost:8080/api/v1/health
   ```

   Harus mengembalikan `{"status":"healthy"}`

2. **Cek Environment Variables:**

   Di `frontend/cms/.env`:

   ```env
   NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1
   NEXTAUTH_SECRET=your-secret-key-change-in-production
   NEXTAUTH_URL=http://localhost:3001
   ```

   Di `backend/.env`:

   ```env
   JWT_SECRET=your-secret-key-change-in-production
   CORS_ALLOWED_ORIGINS=http://localhost:3000,http://localhost:3001
   ```

3. **Cek CORS Configuration:**

   - Pastikan port CMS (biasanya 3001) sudah ditambahkan ke `CORS_ALLOWED_ORIGINS` di backend
   - Restart backend setelah mengubah CORS config

4. **Cek Database:**
   - Pastikan PostgreSQL berjalan
   - Pastikan user admin sudah di-seed: `go run cmd/seed/main.go`

### 2. Invalid Credentials Error

**Kemungkinan Penyebab:**

- Username/password salah
- User belum di-seed
- Password tidak match

**Solusi:**

1. **Cek User di Database:**

   ```sql
   SELECT id, username, email, role, is_active FROM users;
   ```

2. **Re-seed Admin User:**

   ```bash
   cd backend
   go run cmd/seed/main.go
   ```

   Default credentials:

   - Username: `admin`
   - Password: `admin123`

### 3. Network Error / CORS Error

**Kemungkinan Penyebab:**

- Backend tidak accessible dari frontend
- CORS tidak dikonfigurasi dengan benar
- Port berbeda

**Solusi:**

1. **Cek Backend CORS Config:**

   ```env
   # backend/.env
   CORS_ALLOWED_ORIGINS=http://localhost:3000,http://localhost:3001
   ```

2. **Cek API URL di Frontend:**

   ```env
   # frontend/cms/.env
   NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1
   ```

3. **Test API dari Browser Console:**
   ```javascript
   fetch("http://localhost:8080/api/v1/auth/login", {
     method: "POST",
     headers: { "Content-Type": "application/json" },
     body: JSON.stringify({ username: "admin", password: "admin123" }),
   })
     .then((r) => r.json())
     .then(console.log);
   ```

### 4. NextAuth Error

**Kemungkinan Penyebab:**

- NextAuth secret tidak di-set
- NextAuth URL tidak benar
- Route handler error

**Solusi:**

1. **Set NextAuth Secret:**

   ```env
   # frontend/cms/.env
   NEXTAUTH_SECRET=your-secret-key-change-in-production
   NEXTAUTH_URL=http://localhost:3001
   ```

2. **Generate Secret Baru:**

   ```bash
   openssl rand -base64 32
   ```

3. **Cek Browser Console:**
   - Buka Developer Tools (F12)
   - Cek tab Console dan Network untuk error details

## 🐛 Debug Steps

### Step 1: Cek Backend Logs

```bash
cd backend
go run cmd/server/main.go
```

Cari error di console output.

### Step 2: Cek Frontend Logs

```bash
cd frontend/cms
npm run dev
```

Cek browser console untuk error details.

### Step 3: Test API Langsung

```bash
# Test login endpoint
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
    "token_type": "Bearer",
    "expires_in": 900,
    "user": {
      "id": 1,
      "username": "admin",
      "email": "admin@madr.local",
      "role": "admin"
    }
  }
}
```

### Step 4: Cek Database

```bash
# Connect to PostgreSQL
psql -U postgres -d madr_db

# Check users table
SELECT id, username, email, role, is_active FROM users;
```

## ✅ Checklist

- [ ] Backend berjalan di `http://localhost:8080`
- [ ] Frontend CMS berjalan di `http://localhost:3001` (atau port lain)
- [ ] Database PostgreSQL berjalan dan connected
- [ ] User admin sudah di-seed (`go run cmd/seed/main.go`)
- [ ] Environment variables sudah di-set dengan benar
- [ ] CORS_ALLOWED_ORIGINS sudah include port CMS
- [ ] NEXT_PUBLIC_API_URL sudah benar
- [ ] NEXTAUTH_SECRET sudah di-set
- [ ] Browser console tidak ada error CORS

## 📞 Still Having Issues?

1. Cek browser Network tab untuk melihat request/response details
2. Cek backend logs untuk error messages
3. Pastikan semua services (backend, database, frontend) berjalan
4. Coba restart semua services
