# API Documentation - Masjid Management System

Base URL: `http://localhost:8080/api/v1`

## Authentication

> **Note**: Authentication akan diimplementasikan pada tahap selanjutnya. Saat ini endpoint admin belum dilindungi.

## Health Check

### Get Health Status

```http
GET /health
GET /api/health
```

**Response:**
```json
{
  "status": "healthy",
  "service": "Masjid Management System API",
  "version": "1.0.0"
}
```

---

## Announcements

### Get Published Announcements (Public)

Mengambil daftar pengumuman yang sudah dipublish.

```http
GET /announcements/published
```

**Query Parameters:**
- `limit` (optional, default: 10, max: 100) - Jumlah data per halaman
- `offset` (optional, default: 0) - Offset untuk pagination

**Response:**
```json
{
  "data": [
    {
      "id": 1,
      "title": "Pengumuman Sholat Jumat",
      "content": "Sholat Jumat akan dilaksanakan pada pukul 12:00 WIB",
      "is_published": true,
      "published_at": "2024-01-15T10:00:00Z",
      "author": "Admin Masjid",
      "created_at": "2024-01-15T09:00:00Z",
      "updated_at": "2024-01-15T09:00:00Z"
    }
  ],
  "total": 1,
  "limit": 10,
  "offset": 0,
  "total_pages": 1
}
```

**Example:**
```bash
curl http://localhost:8080/api/v1/announcements/published?limit=10&offset=0
```

---

### Get Announcement by ID

Mengambil detail pengumuman berdasarkan ID.

```http
GET /announcements/:id
```

**Path Parameters:**
- `id` (required) - ID pengumuman

**Response:**
```json
{
  "data": {
    "id": 1,
    "title": "Pengumuman Sholat Jumat",
    "content": "Sholat Jumat akan dilaksanakan pada pukul 12:00 WIB",
    "is_published": true,
    "published_at": "2024-01-15T10:00:00Z",
    "author": "Admin Masjid",
    "created_at": "2024-01-15T09:00:00Z",
    "updated_at": "2024-01-15T09:00:00Z"
  }
}
```

**Error Response (404):**
```json
{
  "error": "Announcement not found"
}
```

**Example:**
```bash
curl http://localhost:8080/api/v1/announcements/1
```

---

### Create Announcement (Admin)

Membuat pengumuman baru.

```http
POST /admin/announcements
```

**Request Body:**
```json
{
  "title": "Pengumuman Sholat Jumat",
  "content": "Sholat Jumat akan dilaksanakan pada pukul 12:00 WIB",
  "is_published": true,
  "author": "Admin Masjid"
}
```

**Fields:**
- `title` (required) - Judul pengumuman
- `content` (required) - Isi pengumuman
- `is_published` (optional, default: false) - Status publish
- `author` (optional) - Nama author

**Response (201):**
```json
{
  "message": "Announcement created successfully",
  "data": {
    "id": 1,
    "title": "Pengumuman Sholat Jumat",
    "content": "Sholat Jumat akan dilaksanakan pada pukul 12:00 WIB",
    "is_published": true,
    "published_at": "2024-01-15T10:00:00Z",
    "author": "Admin Masjid",
    "created_at": "2024-01-15T09:00:00Z",
    "updated_at": "2024-01-15T09:00:00Z"
  }
}
```

**Error Response (400):**
```json
{
  "error": "Invalid request body",
  "details": "Key: 'CreateRequest.Title' Error:Field validation for 'Title' failed on the 'required' tag"
}
```

**Example:**
```bash
curl -X POST http://localhost:8080/api/v1/admin/announcements \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Pengumuman Sholat Jumat",
    "content": "Sholat Jumat akan dilaksanakan pada pukul 12:00 WIB",
    "is_published": true,
    "author": "Admin Masjid"
  }'
```

---

### Get All Announcements (Admin)

Mengambil semua pengumuman termasuk yang belum dipublish.

```http
GET /admin/announcements
```

**Query Parameters:**
- `limit` (optional, default: 10, max: 100) - Jumlah data per halaman
- `offset` (optional, default: 0) - Offset untuk pagination

**Response:**
```json
{
  "data": [
    {
      "id": 1,
      "title": "Pengumuman Sholat Jumat",
      "content": "Sholat Jumat akan dilaksanakan pada pukul 12:00 WIB",
      "is_published": true,
      "published_at": "2024-01-15T10:00:00Z",
      "author": "Admin Masjid",
      "created_at": "2024-01-15T09:00:00Z",
      "updated_at": "2024-01-15T09:00:00Z"
    }
  ],
  "total": 1,
  "limit": 10,
  "offset": 0,
  "total_pages": 1
}
```

**Example:**
```bash
curl http://localhost:8080/api/v1/admin/announcements?limit=10&offset=0
```

---

### Update Announcement (Admin)

Mengupdate pengumuman yang sudah ada.

```http
PUT /admin/announcements/:id
```

**Path Parameters:**
- `id` (required) - ID pengumuman

**Request Body:**
```json
{
  "title": "Pengumuman Sholat Jumat (Updated)",
  "content": "Sholat Jumat akan dilaksanakan pada pukul 12:30 WIB",
  "is_published": false,
  "author": "Admin Masjid"
}
```

**Fields (all optional):**
- `title` - Judul pengumuman
- `content` - Isi pengumuman
- `is_published` - Status publish (boolean)
- `author` - Nama author

**Response (200):**
```json
{
  "message": "Announcement updated successfully",
  "data": {
    "id": 1,
    "title": "Pengumuman Sholat Jumat (Updated)",
    "content": "Sholat Jumat akan dilaksanakan pada pukul 12:30 WIB",
    "is_published": false,
    "published_at": null,
    "author": "Admin Masjid",
    "created_at": "2024-01-15T09:00:00Z",
    "updated_at": "2024-01-15T10:30:00Z"
  }
}
```

**Error Response (404):**
```json
{
  "error": "Announcement not found"
}
```

**Example:**
```bash
curl -X PUT http://localhost:8080/api/v1/admin/announcements/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Pengumuman Sholat Jumat (Updated)",
    "is_published": false
  }'
```

---

### Delete Announcement (Admin)

Menghapus pengumuman (soft delete).

```http
DELETE /admin/announcements/:id
```

**Path Parameters:**
- `id` (required) - ID pengumuman

**Response (200):**
```json
{
  "message": "Announcement deleted successfully"
}
```

**Error Response (404):**
```json
{
  "error": "Announcement not found"
}
```

**Example:**
```bash
curl -X DELETE http://localhost:8080/api/v1/admin/announcements/1
```

---

## Error Responses

### 400 Bad Request
```json
{
  "error": "Invalid request body",
  "details": "Error details..."
}
```

### 404 Not Found
```json
{
  "error": "Resource not found"
}
```

### 429 Too Many Requests
```json
{
  "error": "Too many requests. Please try again later."
}
```

### 500 Internal Server Error
```json
{
  "error": "Internal server error"
}
```

---

## Rate Limiting

API dilindungi dengan rate limiting:
- Default: 100 requests per minute per IP
- Dapat dikonfigurasi melalui environment variables

---

## CORS

CORS dikonfigurasi untuk mengizinkan request dari origins yang ditentukan. Default:
- `http://localhost:3000` (Frontend Web)
- `http://localhost:3001` (Frontend CMS)

---

## Notes

- Semua timestamp menggunakan format ISO 8601 (UTC)
- Pagination menggunakan limit/offset pattern
- Soft delete digunakan untuk semua resources
- Admin endpoints akan dilindungi dengan JWT authentication pada tahap selanjutnya

