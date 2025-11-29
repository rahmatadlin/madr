# Database Schema - Masjid Management System

Dokumentasi schema database untuk Masjid Management System.

## Tables

### users

Tabel untuk menyimpan data user/admin.

| Column      | Type          | Constraints           | Description                    |
|-------------|---------------|-----------------------|--------------------------------|
| id          | SERIAL        | PRIMARY KEY           | Auto-increment ID              |
| username    | VARCHAR(100)  | UNIQUE, NOT NULL       | Username unik                  |
| email       | VARCHAR(255)  | UNIQUE, NOT NULL       | Email unik                     |
| password    | VARCHAR(255)  | NOT NULL               | Password (bcrypt hashed)       |
| name        | VARCHAR(255)  |                        | Nama lengkap user              |
| role        | VARCHAR(20)   | DEFAULT 'user'         | Role: 'admin' atau 'user'     |
| is_active   | BOOLEAN       | DEFAULT true           | Status aktif user              |
| last_login  | TIMESTAMP     |                        | Waktu login terakhir           |
| created_at  | TIMESTAMP     | NOT NULL               | Waktu pembuatan                |
| updated_at  | TIMESTAMP     | NOT NULL               | Waktu update terakhir          |
| deleted_at  | TIMESTAMP     |                        | Soft delete timestamp          |

**Indexes:**
- `idx_users_username` (username)
- `idx_users_email` (email)
- `idx_users_deleted_at` (deleted_at)

**Default Admin:**
- Username: `admin`
- Password: `admin123` (harus diubah setelah login pertama)
- Email: `admin@madr.local`
- Role: `admin`

---

### refresh_tokens

Tabel untuk menyimpan refresh token yang dapat di-revoke.

| Column      | Type          | Constraints           | Description                    |
|-------------|---------------|-----------------------|--------------------------------|
| id          | SERIAL        | PRIMARY KEY           | Auto-increment ID              |
| token       | VARCHAR(500)  | UNIQUE, NOT NULL       | Refresh token string           |
| user_id     | INTEGER       | NOT NULL, INDEX        | Foreign key ke users.id        |
| expires_at  | TIMESTAMP     | NOT NULL               | Waktu expired token            |
| is_revoked  | BOOLEAN       | DEFAULT false         | Status revoked token           |
| revoked_at  | TIMESTAMP     |                        | Waktu token di-revoke          |
| user_agent  | VARCHAR(255)  |                        | User agent browser/device      |
| ip_address  | VARCHAR(45)   |                        | IP address user                |
| created_at  | TIMESTAMP     | NOT NULL               | Waktu pembuatan                |
| updated_at  | TIMESTAMP     | NOT NULL               | Waktu update terakhir          |
| deleted_at  | TIMESTAMP     |                        | Soft delete timestamp          |

**Indexes:**
- `idx_refresh_tokens_token` (token)
- `idx_refresh_tokens_user_id` (user_id)
- `idx_refresh_tokens_deleted_at` (deleted_at)

**Features:**
- Token dapat di-revoke untuk logout
- Mendukung logout dari semua device (revoke all tokens)
- Expired tokens dapat dihapus secara berkala

---

### announcements

Tabel untuk menyimpan pengumuman masjid.

| Column       | Type          | Constraints           | Description                    |
|--------------|---------------|-----------------------|--------------------------------|
| id           | SERIAL        | PRIMARY KEY           | Auto-increment ID              |
| title        | VARCHAR(255)  | NOT NULL               | Judul pengumuman               |
| content      | TEXT          | NOT NULL               | Isi pengumuman                 |
| is_published | BOOLEAN       | DEFAULT false         | Status publish                 |
| published_at | TIMESTAMP     |                        | Waktu publish                   |
| author       | VARCHAR(100)  |                        | Nama author                    |
| created_at   | TIMESTAMP     | NOT NULL               | Waktu pembuatan                |
| updated_at   | TIMESTAMP     | NOT NULL               | Waktu update terakhir          |
| deleted_at   | TIMESTAMP     |                        | Soft delete timestamp          |

**Indexes:**
- `idx_announcements_deleted_at` (deleted_at)

---

## Relationships

```
users (1) ──< (N) refresh_tokens
```

- Satu user dapat memiliki banyak refresh tokens
- Refresh token dimiliki oleh satu user

---

## Security Considerations

1. **Password Storage**: Password disimpan menggunakan bcrypt hashing dengan cost 10
2. **Token Storage**: Refresh token disimpan di database untuk mendukung revocation
3. **Soft Delete**: Semua tabel menggunakan soft delete untuk audit trail
4. **Indexes**: Indexes pada kolom yang sering digunakan untuk query optimization

---

## Migration Notes

- Semua tabel menggunakan GORM AutoMigrate untuk development
- Untuk production, gunakan migration tool seperti golang-migrate
- Default admin user dibuat otomatis saat pertama kali aplikasi dijalankan

---

## Future Tables

Tabel-tabel berikut akan ditambahkan pada implementasi selanjutnya:
- `events` - Event dan kegiatan masjid
- `donations` - Data donasi
- `gallery` - Galeri foto
- `banners` - Banner untuk landing page

