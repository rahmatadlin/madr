# YouTube Data API v3 Setup

## Prerequisites

1. **Enable YouTube Data API v3** di Google Cloud Console:
   - Buka: https://console.cloud.google.com/apis/library
   - Cari "YouTube Data API v3"
   - Klik dan pilih **Enable**

2. **Verifikasi API Key**:
   - Buka: https://console.cloud.google.com/apis/credentials
   - Pastikan API key memiliki akses ke **YouTube Data API v3**
   - Jika belum, edit API key → **API restrictions** → pilih "Restrict key" → centang "YouTube Data API v3"

## Konfigurasi

API Key dan Channel ID sudah di-hardcode di:
- `backend/internal/service/youtube/youtube_service.go`

```go
const (
    YouTubeAPIKey    = "AIzaSyAnXmKQn5nka20et5qkyptOySfxmB6h5BY"
    YouTubeChannelID = "UCPTIg2Lw3V81c74awdSBqnQ"
)
```

## Testing

Test endpoint:
```bash
curl http://localhost:8080/api/v1/youtube/kajian
```

Jika error, cek response untuk detail error dari YouTube API.

## Troubleshooting

**Error: "API key not valid"**
- Pastikan YouTube Data API v3 sudah enabled
- Pastikan API key memiliki akses ke YouTube Data API v3

**Error: "The request cannot be completed because you have exceeded your quota"**
- YouTube Data API memiliki quota harian (default 10,000 units)
- Search request = 100 units
- Cek quota di: https://console.cloud.google.com/apis/api/youtube.googleapis.com/quotas

**Error: "Channel not found"**
- Verifikasi Channel ID sudah benar
- Format: `UCPTIg2Lw3V81c74awdSBqnQ` (dimulai dengan UC)
