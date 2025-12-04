# Struktur Migration & Seeding

## 📁 Struktur Folder Migrations

```
backend/
├── migrations/
│   ├── 000001_create_users_table.up.sql          # Migration UP (apply)
│   ├── 000001_create_users_table.down.sql        # Migration DOWN (rollback)
│   ├── 000002_create_refresh_tokens_table.up.sql
│   ├── 000002_create_refresh_tokens_table.down.sql
│   └── ...
```

### Format Naming

`golang-migrate` menggunakan format:
- **`{version}_{name}.up.sql`** - File untuk apply migration (membuat tabel, dll)
- **`{version}_{name}.down.sql`** - File untuk rollback migration (drop tabel, dll)

### Kenapa Tidak Ada Folder `up/` dan `down/`?

**golang-migrate** membaca file langsung dari folder `migrations/` dengan pattern:
- File dengan suffix `.up.sql` → untuk migration UP
- File dengan suffix `.down.sql` → untuk migration DOWN

Jadi **tidak perlu** folder terpisah `up/` dan `down/`. Semua file `.up.sql` dan `.down.sql` berada di folder `migrations/` yang sama.

### Alternatif Struktur (Jika Diperlukan)

Jika ingin menggunakan struktur dengan folder terpisah, bisa menggunakan format:
```
migrations/
├── up/
│   ├── 000001_create_users_table.sql
│   └── 000002_create_refresh_tokens_table.sql
└── down/
    ├── 000001_create_users_table.sql
    └── 000002_create_refresh_tokens_table.sql
```

Tapi ini memerlukan konfigurasi khusus di `migrate.go` dan kurang umum digunakan.

## 🌱 Struktur Seeding

### Folder `seeds/` (Opsional)

Folder `seeds/` bisa digunakan untuk:
- **SQL seed files** - Jika ingin seeding menggunakan SQL langsung
- **JSON/CSV data files** - File data yang akan di-load oleh seeder
- **Scripts** - Script tambahan untuk seeding

### Seeder Saat Ini

Seeder saat ini menggunakan **Go code** di `pkg/seed/seed.go` karena:
- ✅ Lebih fleksibel (bisa validasi, hash password, dll)
- ✅ Type-safe dengan domain models
- ✅ Bisa menggunakan repository layer yang sudah ada
- ✅ Lebih mudah untuk testing

### Contoh Struktur Seeds (Jika Diperlukan)

Jika ingin menggunakan SQL seed files:

```
backend/
├── seeds/
│   ├── 001_admin_users.sql
│   ├── 002_donation_categories.sql
│   └── 003_sample_events.sql
└── pkg/
    └── seed/
        └── seed.go          # Seeder Go code (saat ini)
```

### Kapan Menggunakan SQL Seed Files?

Gunakan SQL seed files jika:
- Data sangat besar (ribuan records)
- Data dari export database lain
- Perlu seeding cepat tanpa validasi Go

Gunakan Go seeder (seperti sekarang) jika:
- Perlu validasi dan transformasi data
- Perlu menggunakan business logic (hash password, dll)
- Perlu idempotent checking (skip jika sudah ada)

## 📝 Kesimpulan

1. **Folder `migrations/up/` dan `migrations/down/`** → **TIDAK DIPERLUKAN** (sudah dihapus)
   - File migration langsung di `migrations/` dengan suffix `.up.sql` dan `.down.sql`

2. **Folder `seeds/`** → **OPSIONAL** (saat ini kosong)
   - Bisa digunakan untuk SQL seed files di masa depan
   - Saat ini seeder menggunakan Go code di `pkg/seed/seed.go`

3. **Struktur saat ini sudah benar** untuk `golang-migrate` standard format.

