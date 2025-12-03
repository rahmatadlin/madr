# API Documentation - Masjid Management System

Base URL: `http://localhost:8080/api/v1`

## Authentication

API menggunakan JWT (JSON Web Token) untuk authentication. Terdapat dua jenis token:

- **Access Token**: Token untuk mengakses protected endpoints (expires in 15 minutes by default)
- **Refresh Token**: Token untuk memperbarui access token (expires in 7 days by default)

### Authentication Flow

1. **Login** → Dapatkan access token dan refresh token
2. **Gunakan Access Token** → Setiap request ke protected endpoint harus menyertakan header: `Authorization: Bearer <access_token>`
3. **Refresh Token** → Ketika access token expired, gunakan refresh token untuk mendapatkan access token baru
4. **Logout** → Revoke refresh token untuk logout

### Default Admin Credentials

Setelah pertama kali menjalankan aplikasi, default admin akan dibuat secara otomatis:

- **Username**: `admin`
- **Password**: `admin123`
- **Email**: `admin@madr.local`

> **⚠️ PENTING**: Ubah password default admin setelah pertama kali login!

---

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

## Events

### Get All Events (Public)

Mengambil daftar semua event yang tersedia.

```http
GET /events
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
      "title": "Sholat Jumat Berjamaah",
      "description": "Sholat Jumat akan dilaksanakan di masjid utama",
      "date": "2024-01-20T12:00:00Z",
      "location": "Masjid Al-Madr",
      "created_at": "2024-01-15T10:00:00Z",
      "updated_at": "2024-01-15T10:00:00Z"
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
curl http://localhost:8080/api/v1/events?limit=10&offset=0
```

---

### Get Event by ID (Public)

Mengambil detail event berdasarkan ID.

```http
GET /events/:id
```

**Path Parameters:**

- `id` (required) - ID event

**Response:**

```json
{
  "data": {
    "id": 1,
    "title": "Sholat Jumat Berjamaah",
    "description": "Sholat Jumat akan dilaksanakan di masjid utama",
    "date": "2024-01-20T12:00:00Z",
    "location": "Masjid Al-Madr",
    "created_at": "2024-01-15T10:00:00Z",
    "updated_at": "2024-01-15T10:00:00Z"
  }
}
```

**Example:**

```bash
curl http://localhost:8080/api/v1/events/1
```

---

### Create Event (Admin - Protected)

Membuat event baru.

```http
POST /admin/events
```

**Headers:**

```
Authorization: Bearer <access_token>
```

**Request Body:**

```json
{
  "title": "Sholat Jumat Berjamaah",
  "description": "Sholat Jumat akan dilaksanakan di masjid utama",
  "date": "2024-01-20T12:00:00Z",
  "location": "Masjid Al-Madr"
}
```

**Fields:**

- `title` (required, min: 3, max: 255) - Judul event
- `description` (optional) - Deskripsi event
- `date` (required) - Tanggal dan waktu event (ISO 8601 format)
- `location` (optional, max: 255) - Lokasi event

**Response (201):**

```json
{
  "message": "Event created successfully",
  "data": {
    "id": 1,
    "title": "Sholat Jumat Berjamaah",
    "description": "Sholat Jumat akan dilaksanakan di masjid utama",
    "date": "2024-01-20T12:00:00Z",
    "location": "Masjid Al-Madr",
    "created_at": "2024-01-15T10:00:00Z",
    "updated_at": "2024-01-15T10:00:00Z"
  }
}
```

**Example:**

```bash
curl -X POST http://localhost:8080/api/v1/admin/events \
  -H "Authorization: Bearer <access_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Sholat Jumat Berjamaah",
    "description": "Sholat Jumat akan dilaksanakan di masjid utama",
    "date": "2024-01-20T12:00:00Z",
    "location": "Masjid Al-Madr"
  }'
```

---

### Update Event (Admin - Protected)

Mengupdate event yang sudah ada.

```http
PUT /admin/events/:id
```

**Headers:**

```
Authorization: Bearer <access_token>
```

**Path Parameters:**

- `id` (required) - ID event

**Request Body:**

```json
{
  "title": "Sholat Jumat Berjamaah (Updated)",
  "description": "Updated description",
  "date": "2024-01-20T13:00:00Z",
  "location": "Masjid Al-Madr - Ruang Utama"
}
```

**Fields (all optional):**

- `title` - Judul event
- `description` - Deskripsi event
- `date` - Tanggal dan waktu event
- `location` - Lokasi event

**Response (200):**

```json
{
  "message": "Event updated successfully",
  "data": {
    "id": 1,
    "title": "Sholat Jumat Berjamaah (Updated)",
    "description": "Updated description",
    "date": "2024-01-20T13:00:00Z",
    "location": "Masjid Al-Madr - Ruang Utama",
    "created_at": "2024-01-15T10:00:00Z",
    "updated_at": "2024-01-15T11:00:00Z"
  }
}
```

**Example:**

```bash
curl -X PUT http://localhost:8080/api/v1/admin/events/1 \
  -H "Authorization: Bearer <access_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Sholat Jumat Berjamaah (Updated)",
    "location": "Masjid Al-Madr - Ruang Utama"
  }'
```

---

### Delete Event (Admin - Protected)

Menghapus event (soft delete).

```http
DELETE /admin/events/:id
```

**Headers:**

```
Authorization: Bearer <access_token>
```

**Path Parameters:**

- `id` (required) - ID event

**Response (200):**

```json
{
  "message": "Event deleted successfully"
}
```

**Example:**

```bash
curl -X DELETE http://localhost:8080/api/v1/admin/events/1 \
  -H "Authorization: Bearer <access_token>"
```

---

## Gallery

### Get All Gallery Items (Public)

Mengambil daftar semua item galeri.

```http
GET /gallery
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
      "title": "Foto Kegiatan Sholat Jumat",
      "image_url": "uploads/gallery/foto-jumat-2024.jpg",
      "created_at": "2024-01-15T10:00:00Z",
      "updated_at": "2024-01-15T10:00:00Z"
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
curl http://localhost:8080/api/v1/gallery?limit=10&offset=0
```

---

### Create Gallery Item (Admin - Protected)

Menambahkan item baru ke galeri dengan file upload.

```http
POST /admin/gallery
```

**Headers:**

```
Authorization: Bearer <access_token>
Content-Type: multipart/form-data
```

**Form Data:**

- `file` (required) - File gambar (jpg, jpeg, png, webp)
- `title` (required, min: 3, max: 255) - Judul foto

**File Requirements:**

- Max size: 10MB (configurable)
- Allowed types: image/jpeg, image/jpg, image/png, image/webp

**Response (201):**

```json
{
  "message": "Gallery item created successfully",
  "data": {
    "id": 1,
    "title": "Foto Kegiatan Sholat Jumat",
    "image_url": "http://localhost:8080/uploads/1705312800_abc123_foto_jumat.jpg",
    "created_at": "2024-01-15T10:00:00Z",
    "updated_at": "2024-01-15T10:00:00Z"
  }
}
```

**Example with cURL:**

```bash
curl -X POST http://localhost:8080/api/v1/admin/gallery \
  -H "Authorization: Bearer <access_token>" \
  -F "file=@/path/to/image.jpg" \
  -F "title=Foto Kegiatan Sholat Jumat"
```

**Example with JavaScript (FormData):**

```javascript
const formData = new FormData();
formData.append("file", fileInput.files[0]);
formData.append("title", "Foto Kegiatan Sholat Jumat");

fetch("http://localhost:8080/api/v1/admin/gallery", {
  method: "POST",
  headers: {
    Authorization: "Bearer <access_token>",
  },
  body: formData,
});
```

---

### Delete Gallery Item (Admin - Protected)

Menghapus item dari galeri (soft delete).

```http
DELETE /admin/gallery/:id
```

**Headers:**

```
Authorization: Bearer <access_token>
```

**Path Parameters:**

- `id` (required) - ID gallery item

**Response (200):**

```json
{
  "message": "Gallery item deleted successfully"
}
```

**Example:**

```bash
curl -X DELETE http://localhost:8080/api/v1/admin/gallery/1 \
  -H "Authorization: Bearer <access_token>"
```

---

## Banners

### Get All Banners (Public)

Mengambil daftar semua banner.

```http
GET /banners
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
      "title": "Banner Sholat Jumat",
      "media_url": "uploads/banners/banner-jumat.jpg",
      "type": "image",
      "created_at": "2024-01-15T10:00:00Z",
      "updated_at": "2024-01-15T10:00:00Z"
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
curl http://localhost:8080/api/v1/banners?limit=10&offset=0
```

---

### Get Banner by ID (Public)

Mengambil detail banner berdasarkan ID.

```http
GET /banners/:id
```

**Path Parameters:**

- `id` (required) - ID banner

**Response:**

```json
{
  "data": {
    "id": 1,
    "title": "Banner Sholat Jumat",
    "media_url": "uploads/banners/banner-jumat.jpg",
    "type": "image",
    "created_at": "2024-01-15T10:00:00Z",
    "updated_at": "2024-01-15T10:00:00Z"
  }
}
```

**Example:**

```bash
curl http://localhost:8080/api/v1/banners/1
```

---

### Create Banner (Admin - Protected)

Membuat banner baru dengan file upload.

```http
POST /admin/banners
```

**Headers:**

```
Authorization: Bearer <access_token>
Content-Type: multipart/form-data
```

**Form Data:**

- `file` (required) - File media (image: jpg, jpeg, png, webp atau video: mp4)
- `title` (required, min: 3, max: 255) - Judul banner
- `type` (required) - Tipe media: `"image"` atau `"video"`

**File Requirements:**

- Max size: 10MB (configurable)
- Image types: image/jpeg, image/jpg, image/png, image/webp
- Video types: video/mp4
- File type must match the `type` field

**Response (201):**

```json
{
  "message": "Banner created successfully",
  "data": {
    "id": 1,
    "title": "Banner Sholat Jumat",
    "media_url": "http://localhost:8080/uploads/1705312800_abc123_banner_jumat.jpg",
    "type": "image",
    "created_at": "2024-01-15T10:00:00Z",
    "updated_at": "2024-01-15T10:00:00Z"
  }
}
```

**Example with cURL (Image):**

```bash
curl -X POST http://localhost:8080/api/v1/admin/banners \
  -H "Authorization: Bearer <access_token>" \
  -F "file=@/path/to/banner.jpg" \
  -F "title=Banner Sholat Jumat" \
  -F "type=image"
```

**Example with cURL (Video):**

```bash
curl -X POST http://localhost:8080/api/v1/admin/banners \
  -H "Authorization: Bearer <access_token>" \
  -F "file=@/path/to/video.mp4" \
  -F "title=Video Pengumuman" \
  -F "type=video"
```

**Example with JavaScript (FormData):**

```javascript
const formData = new FormData();
formData.append("file", fileInput.files[0]);
formData.append("title", "Banner Sholat Jumat");
formData.append("type", "image");

fetch("http://localhost:8080/api/v1/admin/banners", {
  method: "POST",
  headers: {
    Authorization: "Bearer <access_token>",
  },
  body: formData,
});
```

---

### Update Banner (Admin - Protected)

Mengupdate banner yang sudah ada. File upload bersifat optional.

```http
PUT /admin/banners/:id
```

**Headers:**

```
Authorization: Bearer <access_token>
Content-Type: multipart/form-data
```

**Path Parameters:**

- `id` (required) - ID banner

**Form Data (all optional):**

- `file` - File media baru (jika ingin mengganti media)
- `title` - Judul banner
- `type` - Tipe media: `"image"` atau `"video"` (required jika upload file baru)

**File Requirements (jika upload file baru):**

- Max size: 10MB (configurable)
- Image types: image/jpeg, image/jpg, image/png, image/webp
- Video types: video/mp4
- File type must match the `type` field

**Response (200):**

```json
{
  "message": "Banner updated successfully",
  "data": {
    "id": 1,
    "title": "Banner Sholat Jumat (Updated)",
    "media_url": "http://localhost:8080/uploads/1705312800_abc123_banner_jumat_new.jpg",
    "type": "image",
    "created_at": "2024-01-15T10:00:00Z",
    "updated_at": "2024-01-15T11:00:00Z"
  }
}
```

**Example with cURL (Update title only):**

```bash
curl -X PUT http://localhost:8080/api/v1/admin/banners/1 \
  -H "Authorization: Bearer <access_token>" \
  -F "title=Banner Sholat Jumat (Updated)"
```

**Example with cURL (Update with new file):**

```bash
curl -X PUT http://localhost:8080/api/v1/admin/banners/1 \
  -H "Authorization: Bearer <access_token>" \
  -F "file=@/path/to/new-banner.jpg" \
  -F "title=Banner Sholat Jumat (Updated)" \
  -F "type=image"
```

**Example with JavaScript (FormData):**

```javascript
const formData = new FormData();
formData.append("title", "Banner Sholat Jumat (Updated)");
// Optional: upload new file
if (fileInput.files[0]) {
  formData.append("file", fileInput.files[0]);
  formData.append("type", "image");
}

fetch("http://localhost:8080/api/v1/admin/banners/1", {
  method: "PUT",
  headers: {
    Authorization: "Bearer <access_token>",
  },
  body: formData,
});
```

---

### Delete Banner (Admin - Protected)

Menghapus banner (soft delete).

```http
DELETE /admin/banners/:id
```

**Headers:**

```
Authorization: Bearer <access_token>
```

**Path Parameters:**

- `id` (required) - ID banner

**Response (200):**

```json
{
  "message": "Banner deleted successfully"
}
```

**Example:**

```bash
curl -X DELETE http://localhost:8080/api/v1/admin/banners/1 \
  -H "Authorization: Bearer <access_token>"
```

---

## File Upload Endpoint

### Upload File (Admin - Protected)

Utility endpoint untuk upload file secara standalone (optional, bisa langsung upload via Gallery/Banner endpoints).

```http
POST /admin/upload
```

**Headers:**

```
Authorization: Bearer <access_token>
Content-Type: multipart/form-data
```

**Form Data:**

- `file` (required) - File yang akan diupload

**File Requirements:**

- Max size: 10MB (default, configurable via `UPLOAD_MAX_SIZE`)
- Allowed types: image/jpeg, image/jpg, image/png, image/webp, video/mp4

**Response (200):**

```json
{
  "message": "File uploaded successfully",
  "data": {
    "filename": "1705312800_abc123_original_name.jpg",
    "original_name": "original_name.jpg",
    "url": "http://localhost:8080/uploads/1705312800_abc123_original_name.jpg",
    "mime_type": "image/jpeg",
    "size": 245678
  }
}
```

**Example with cURL:**

```bash
curl -X POST http://localhost:8080/api/v1/admin/upload \
  -H "Authorization: Bearer <access_token>" \
  -F "file=@/path/to/file.jpg"
```

**Error Responses:**

**400 Bad Request - File size exceeds limit:**

```json
{
  "error": "File size exceeds maximum allowed size (10485760 bytes)"
}
```

**400 Bad Request - Invalid file type:**

```json
{
  "error": "File type 'application/pdf' is not allowed"
}
```

**400 Bad Request - File is required:**

```json
{
  "error": "File is required"
}
```

---

## Static File Serving

Uploaded files dapat diakses secara langsung melalui:

```
GET /uploads/<filename>
```

**Example:**

```bash
curl http://localhost:8080/uploads/1705312800_abc123_image.jpg
```

Files disimpan di folder `./uploads` (configurable via `UPLOAD_PATH`).

---

## File Upload Configuration

File upload dapat dikonfigurasi melalui environment variables:

- `UPLOAD_MAX_SIZE` - Maximum file size in bytes (default: 10485760 = 10MB)
- `UPLOAD_ALLOWED_TYPES` - Comma-separated list of allowed MIME types
- `UPLOAD_PATH` - Path to save uploaded files (default: `./uploads`)
- `UPLOAD_PUBLIC_URL` - Public URL prefix for uploaded files (default: `http://localhost:8080/uploads`)

**Example .env configuration:**

```env
UPLOAD_MAX_SIZE=10485760
UPLOAD_ALLOWED_TYPES=image/jpeg,image/jpg,image/png,image/webp,video/mp4
UPLOAD_PATH=./uploads
UPLOAD_PUBLIC_URL=http://localhost:8080/uploads
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

## Authentication Endpoints

### Register User

Mendaftarkan user baru.

```http
POST /auth/register
```

**Request Body:**

```json
{
  "username": "newuser",
  "email": "user@example.com",
  "password": "password123",
  "name": "New User",
  "role": "user"
}
```

**Fields:**

- `username` (required, min: 3, max: 100) - Username unik
- `email` (required, valid email) - Email unik
- `password` (required, min: 6) - Password
- `name` (optional, max: 255) - Nama lengkap
- `role` (optional) - Role user: "user" (default) atau "admin"

**Response (201):**

```json
{
  "message": "User registered successfully",
  "data": {
    "user": {
      "id": 1,
      "username": "newuser",
      "email": "user@example.com",
      "name": "New User",
      "role": "user",
      "is_active": true,
      "created_at": "2024-01-15T10:00:00Z",
      "updated_at": "2024-01-15T10:00:00Z"
    }
  }
}
```

**Error Response (409):**

```json
{
  "error": "username already exists"
}
```

**Example:**

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "newuser",
    "email": "user@example.com",
    "password": "password123",
    "name": "New User"
  }'
```

---

### Login

Login dan mendapatkan access token serta refresh token.

```http
POST /auth/login
```

**Request Body:**

```json
{
  "username": "admin",
  "password": "admin123"
}
```

**Fields:**

- `username` (required) - Username atau email
- `password` (required) - Password

**Response (200):**

```json
{
  "message": "Login successful",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "token_type": "Bearer",
    "expires_in": 900,
    "user": {
      "id": 1,
      "username": "admin",
      "email": "admin@madr.local",
      "name": "Default Admin",
      "role": "admin",
      "is_active": true,
      "last_login": "2024-01-15T10:00:00Z"
    }
  }
}
```

**Error Response (401):**

```json
{
  "error": "Invalid credentials"
}
```

**Error Response (403):**

```json
{
  "error": "Account is inactive"
}
```

**Example:**

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "admin123"
  }'
```

---

### Refresh Token

Memperbarui access token menggunakan refresh token.

```http
POST /auth/refresh
```

**Request Body:**

```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Fields:**

- `refresh_token` (required) - Refresh token yang valid

**Response (200):**

```json
{
  "message": "Token refreshed successfully",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "token_type": "Bearer",
    "expires_in": 900
  }
}
```

**Error Response (401):**

```json
{
  "error": "invalid refresh token"
}
```

**Example:**

```bash
curl -X POST http://localhost:8080/api/v1/auth/refresh \
  -H "Content-Type: application/json" \
  -d '{
    "refresh_token": "your-refresh-token-here"
  }'
```

---

### Get Current User (Protected)

Mendapatkan informasi user yang sedang login.

```http
GET /auth/me
```

**Headers:**

```
Authorization: Bearer <access_token>
```

**Response (200):**

```json
{
  "data": {
    "user": {
      "id": 1,
      "username": "admin",
      "email": "admin@madr.local",
      "name": "Default Admin",
      "role": "admin",
      "is_active": true,
      "last_login": "2024-01-15T10:00:00Z",
      "created_at": "2024-01-15T09:00:00Z",
      "updated_at": "2024-01-15T10:00:00Z"
    }
  }
}
```

**Error Response (401):**

```json
{
  "error": "Unauthorized"
}
```

**Example:**

```bash
curl http://localhost:8080/api/v1/auth/me \
  -H "Authorization: Bearer your-access-token-here"
```

---

### Logout

Logout dan revoke refresh token.

```http
POST /auth/logout
```

**Request Body:**

```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Fields:**

- `refresh_token` (required) - Refresh token yang akan di-revoke

**Response (200):**

```json
{
  "message": "Logged out successfully"
}
```

**Example:**

```bash
curl -X POST http://localhost:8080/api/v1/auth/logout \
  -H "Content-Type: application/json" \
  -d '{
    "refresh_token": "your-refresh-token-here"
  }'
```

---

### Logout All Devices (Protected)

Logout dari semua device dengan me-revoke semua refresh token user.

```http
POST /auth/logout-all
```

**Headers:**

```
Authorization: Bearer <access_token>
```

**Response (200):**

```json
{
  "message": "Logged out from all devices successfully"
}
```

**Error Response (401):**

```json
{
  "error": "Unauthorized"
}
```

**Example:**

```bash
curl -X POST http://localhost:8080/api/v1/auth/logout-all \
  -H "Authorization: Bearer your-access-token-here"
```

---

## Notes

- Semua timestamp menggunakan format ISO 8601 (UTC)
- Pagination menggunakan limit/offset pattern
- Soft delete digunakan untuk semua resources
- **Admin endpoints sekarang dilindungi dengan JWT authentication**
- Access token expired dalam 15 menit (default)
- Refresh token expired dalam 7 hari (default)
- Refresh token disimpan di database untuk mendukung revocation
- Password di-hash menggunakan bcrypt

---

## Summary of Endpoints

### Public Endpoints (No Authentication Required)

- `GET /announcements/published` - Get published announcements
- `GET /announcements/:id` - Get announcement by ID
- `GET /events` - Get all events
- `GET /events/:id` - Get event by ID
- `GET /gallery` - Get all gallery items
- `GET /banners` - Get all banners
- `GET /banners/:id` - Get banner by ID
- `POST /auth/register` - Register new user
- `POST /auth/login` - Login
- `POST /auth/refresh` - Refresh access token
- `POST /auth/logout` - Logout

### Protected Endpoints (Require JWT Authentication)

- `GET /auth/me` - Get current user info
- `POST /auth/logout-all` - Logout from all devices

### Admin Endpoints (Require JWT Authentication + Admin Role)

- `POST /admin/announcements` - Create announcement
- `GET /admin/announcements` - Get all announcements (including unpublished)
- `PUT /admin/announcements/:id` - Update announcement
- `DELETE /admin/announcements/:id` - Delete announcement
- `POST /admin/events` - Create event
- `PUT /admin/events/:id` - Update event
- `DELETE /admin/events/:id` - Delete event
- `POST /admin/gallery` - Create gallery item
- `DELETE /admin/gallery/:id` - Delete gallery item
- `POST /admin/banners` - Create banner
- `PUT /admin/banners/:id` - Update banner
- `DELETE /admin/banners/:id` - Delete banner

---

## Donation Categories

### Get All Donation Categories (Admin - Protected)

Mengambil daftar semua kategori donasi.

```http
GET /admin/donation-categories
```

**Headers:**

```
Authorization: Bearer <access_token>
```

**Response:**

```json
{
  "data": [
    {
      "id": 1,
      "name": "Pembangunan",
      "description": "Donasi untuk pembangunan masjid",
      "created_at": "2024-01-15T10:00:00Z",
      "updated_at": "2024-01-15T10:00:00Z"
    },
    {
      "id": 2,
      "name": "Operasional",
      "description": "Donasi untuk operasional masjid",
      "created_at": "2024-01-15T10:00:00Z",
      "updated_at": "2024-01-15T10:00:00Z"
    }
  ]
}
```

**Example:**

```bash
curl http://localhost:8080/api/v1/admin/donation-categories \
  -H "Authorization: Bearer <access_token>"
```

---

### Get Donation Category by ID (Admin - Protected)

Mengambil detail kategori donasi berdasarkan ID.

```http
GET /admin/donation-categories/:id
```

**Headers:**

```
Authorization: Bearer <access_token>
```

**Response:**

```json
{
  "data": {
    "id": 1,
    "name": "Pembangunan",
    "description": "Donasi untuk pembangunan masjid",
    "created_at": "2024-01-15T10:00:00Z",
    "updated_at": "2024-01-15T10:00:00Z"
  }
}
```

**Example:**

```bash
curl http://localhost:8080/api/v1/admin/donation-categories/1 \
  -H "Authorization: Bearer <access_token>"
```

---

### Create Donation Category (Admin - Protected)

Membuat kategori donasi baru.

```http
POST /admin/donation-categories
```

**Headers:**

```
Authorization: Bearer <access_token>
Content-Type: application/json
```

**Request Body:**

```json
{
  "name": "Pembangunan",
  "description": "Donasi untuk pembangunan masjid"
}
```

**Fields:**

- `name` (required, min: 3, max: 255) - Nama kategori (unique)
- `description` (optional) - Deskripsi kategori

**Response (201):**

```json
{
  "message": "Donation category created successfully",
  "data": {
    "id": 1,
    "name": "Pembangunan",
    "description": "Donasi untuk pembangunan masjid",
    "created_at": "2024-01-15T10:00:00Z",
    "updated_at": "2024-01-15T10:00:00Z"
  }
}
```

**Error Response (409):**

```json
{
  "error": "category name already exists"
}
```

**Example:**

```bash
curl -X POST http://localhost:8080/api/v1/admin/donation-categories \
  -H "Authorization: Bearer <access_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Pembangunan",
    "description": "Donasi untuk pembangunan masjid"
  }'
```

---

### Update Donation Category (Admin - Protected)

Mengupdate kategori donasi yang sudah ada.

```http
PUT /admin/donation-categories/:id
```

**Headers:**

```
Authorization: Bearer <access_token>
Content-Type: application/json
```

**Request Body:**

```json
{
  "name": "Pembangunan (Updated)",
  "description": "Updated description"
}
```

**Fields (all optional):**

- `name` - Nama kategori
- `description` - Deskripsi kategori

**Response (200):**

```json
{
  "message": "Donation category updated successfully",
  "data": {
    "id": 1,
    "name": "Pembangunan (Updated)",
    "description": "Updated description",
    "created_at": "2024-01-15T10:00:00Z",
    "updated_at": "2024-01-15T11:00:00Z"
  }
}
```

**Example:**

```bash
curl -X PUT http://localhost:8080/api/v1/admin/donation-categories/1 \
  -H "Authorization: Bearer <access_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Pembangunan (Updated)"
  }'
```

---

### Delete Donation Category (Admin - Protected)

Menghapus kategori donasi (soft delete).

```http
DELETE /admin/donation-categories/:id
```

**Headers:**

```
Authorization: Bearer <access_token>
```

**Response (200):**

```json
{
  "message": "Donation category deleted successfully"
}
```

**Example:**

```bash
curl -X DELETE http://localhost:8080/api/v1/admin/donation-categories/1 \
  -H "Authorization: Bearer <access_token>"
```

---

## Donations

### Get All Donations (Admin - Protected)

Mengambil daftar semua donasi dengan pagination dan optional status filter.

```http
GET /admin/donations
```

**Headers:**

```
Authorization: Bearer <access_token>
```

**Query Parameters:**

- `limit` (optional, default: 10, max: 100) - Jumlah data per halaman
- `offset` (optional, default: 0) - Offset untuk pagination
- `status` (optional) - Filter by payment status: `pending`, `success`, `failed`

**Response:**

```json
{
  "data": [
    {
      "id": 1,
      "category_id": 1,
      "donor_name": "Ahmad Fauzi",
      "amount": 100000,
      "message": "Semoga berkah",
      "payment_status": "success",
      "category": {
        "id": 1,
        "name": "Pembangunan",
        "description": "Donasi untuk pembangunan masjid"
      },
      "created_at": "2024-01-15T10:00:00Z",
      "updated_at": "2024-01-15T10:00:00Z"
    }
  ],
  "total": 150,
  "limit": 10,
  "offset": 0,
  "total_pages": 15
}
```

**Example:**

```bash
# Get all donations
curl http://localhost:8080/api/v1/admin/donations?limit=10&offset=0 \
  -H "Authorization: Bearer <access_token>"

# Get only success donations
curl http://localhost:8080/api/v1/admin/donations?status=success \
  -H "Authorization: Bearer <access_token>"
```

---

### Get Donation by ID (Admin - Protected)

Mengambil detail donasi berdasarkan ID.

```http
GET /admin/donations/:id
```

**Headers:**

```
Authorization: Bearer <access_token>
```

**Response:**

```json
{
  "data": {
    "id": 1,
    "category_id": 1,
    "donor_name": "Ahmad Fauzi",
    "amount": 100000,
    "message": "Semoga berkah",
    "payment_status": "success",
    "category": {
      "id": 1,
      "name": "Pembangunan",
      "description": "Donasi untuk pembangunan masjid"
    },
    "created_at": "2024-01-15T10:00:00Z",
    "updated_at": "2024-01-15T10:00:00Z"
  }
}
```

**Example:**

```bash
curl http://localhost:8080/api/v1/admin/donations/1 \
  -H "Authorization: Bearer <access_token>"
```

---

### Create Donation (Admin - Protected)

Membuat entri donasi baru (manual entry dari admin).

```http
POST /admin/donations
```

**Headers:**

```
Authorization: Bearer <access_token>
Content-Type: application/json
```

**Request Body:**

```json
{
  "category_id": 1,
  "donor_name": "Ahmad Fauzi",
  "amount": 100000,
  "message": "Semoga berkah",
  "payment_status": "success"
}
```

**Fields:**

- `category_id` (required) - ID kategori donasi
- `donor_name` (optional) - Nama donor (null untuk anonymous)
- `amount` (required, must be > 0) - Jumlah donasi
- `message` (optional) - Pesan dari donor
- `payment_status` (optional) - Status pembayaran: `pending` (default), `success`, `failed`

**Response (201):**

```json
{
  "message": "Donation created successfully",
  "data": {
    "id": 1,
    "category_id": 1,
    "donor_name": "Ahmad Fauzi",
    "amount": 100000,
    "message": "Semoga berkah",
    "payment_status": "success",
    "created_at": "2024-01-15T10:00:00Z",
    "updated_at": "2024-01-15T10:00:00Z"
  }
}
```

**Example:**

```bash
curl -X POST http://localhost:8080/api/v1/admin/donations \
  -H "Authorization: Bearer <access_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "category_id": 1,
    "donor_name": "Ahmad Fauzi",
    "amount": 100000,
    "message": "Semoga berkah",
    "payment_status": "success"
  }'
```

**Example - Anonymous Donation:**

```bash
curl -X POST http://localhost:8080/api/v1/admin/donations \
  -H "Authorization: Bearer <access_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "category_id": 1,
    "amount": 50000,
    "payment_status": "success"
  }'
```

---

### Update Donation (Admin - Protected)

Mengupdate donasi yang sudah ada.

```http
PUT /admin/donations/:id
```

**Headers:**

```
Authorization: Bearer <access_token>
Content-Type: application/json
```

**Request Body:**

```json
{
  "category_id": 2,
  "donor_name": "Ahmad Fauzi (Updated)",
  "amount": 150000,
  "message": "Updated message",
  "payment_status": "success"
}
```

**Fields (all optional):**

- `category_id` - ID kategori donasi
- `donor_name` - Nama donor (null untuk anonymous)
- `amount` - Jumlah donasi (must be > 0)
- `message` - Pesan dari donor
- `payment_status` - Status pembayaran: `pending`, `success`, `failed`

**Response (200):**

```json
{
  "message": "Donation updated successfully",
  "data": {
    "id": 1,
    "category_id": 2,
    "donor_name": "Ahmad Fauzi (Updated)",
    "amount": 150000,
    "message": "Updated message",
    "payment_status": "success",
    "created_at": "2024-01-15T10:00:00Z",
    "updated_at": "2024-01-15T11:00:00Z"
  }
}
```

**Example:**

```bash
curl -X PUT http://localhost:8080/api/v1/admin/donations/1 \
  -H "Authorization: Bearer <access_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "amount": 150000,
    "payment_status": "success"
  }'
```

---

### Delete Donation (Admin - Protected)

Menghapus donasi (soft delete).

```http
DELETE /admin/donations/:id
```

**Headers:**

```
Authorization: Bearer <access_token>
```

**Response (200):**

```json
{
  "message": "Donation deleted successfully"
}
```

**Example:**

```bash
curl -X DELETE http://localhost:8080/api/v1/admin/donations/1 \
  -H "Authorization: Bearer <access_token>"
```

---

### Get Donation Summary (Public)

Mendapatkan ringkasan donasi untuk ditampilkan di landing page. **Endpoint ini PUBLIC** (tidak memerlukan authentication).

```http
GET /donations/summary
```

**Response:**

```json
{
  "data": {
    "total_amount": 20000000,
    "total_transactions": 150,
    "per_category": [
      {
        "category_id": 1,
        "category": "Pembangunan",
        "amount": 12000000
      },
      {
        "category_id": 2,
        "category": "Operasional",
        "amount": 6000000
      },
      {
        "category_id": 3,
        "category": "Sosial",
        "amount": 2000000
      }
    ]
  }
}
```

**Notes:**

- Summary hanya menghitung donasi dengan status `success`
- Total amount dan transactions dihitung dari semua donasi success
- Per category diurutkan berdasarkan amount descending
- Menggunakan optimized SQL dengan JOIN dan GROUP BY

**Example:**

```bash
curl http://localhost:8080/api/v1/donations/summary
```

---

## Next Improvements Suggestions

### 1. File Upload Endpoint

Implementasi endpoint untuk upload file (gambar/video) untuk Gallery dan Banner:

- `POST /admin/upload` - Upload file dan return URL
- Support multiple file formats (jpg, png, mp4, etc.)
- File validation (size, type)
- Storage di `/uploads` folder atau cloud storage

### 2. Advanced Filtering & Search

- Filter events by date range
- Search announcements by keyword
- Filter gallery by category/tags
- Sort options untuk semua endpoints

### 3. Donations Module

- CRUD untuk donasi
- Payment gateway integration
- Donation reports & analytics
- Recurring donations support

### 4. Enhanced Features

- Event registration/RSVP
- Comments system untuk announcements
- Newsletter subscription
- Email notifications untuk events
- SMS notifications

### 5. Performance & Optimization

- Caching untuk public endpoints (Redis)
- Image optimization & CDN integration
- Database indexing optimization
- API response compression

### 6. Security Enhancements

- Rate limiting per user (not just IP)
- API key for public integrations
- Audit logging untuk admin actions
- Two-factor authentication (2FA)

### 7. Analytics & Reporting

- Dashboard analytics
- User activity tracking
- Content performance metrics
- Export reports (PDF/Excel)

### 8. Multi-language Support

- i18n untuk announcements & events
- Language preference per user
- Content translation management
