# ✅ NextAuth 500 Error - FIXED!

## 🔍 Masalah yang Terjadi

Error `TypeError: Function.prototype.apply was called on #<Object>, which is an object and not a function` terjadi karena NextAuth v5 beta mengembalikan object dengan method `GET` dan `POST`, bukan langsung function.

## ✅ Solusi

Route handler sudah diperbaiki untuk NextAuth v5 beta:

```typescript
import NextAuth from "next-auth";
import { authOptions } from "@/lib/auth";

// NextAuth v5 beta - NextAuth returns handlers object with GET and POST methods
const { handlers } = NextAuth(authOptions);

// Export GET and POST handlers
export const GET = handlers.GET;
export const POST = handlers.POST;
```

## 🧪 Test

Test endpoint session:

```bash
curl http://localhost:3000/api/auth/session
```

Expected response (belum login):
```json
null
```

Expected response (sudah login):
```json
{
  "user": {
    "id": "...",
    "name": "...",
    "email": "...",
    "role": "..."
  },
  "expires": "..."
}
```

## ✅ Status

- [x] Route handler fixed
- [x] Session endpoint working (returns `null` when not logged in)
- [x] No more 500 errors
- [x] Ready for login testing

## 🚀 Next Steps

1. **Test Login** - Coba login dari browser di `http://localhost:3000/login`
2. **Check Session** - Setelah login, cek session endpoint untuk melihat user data
3. **Test Protected Routes** - Coba akses dashboard untuk test protected routes

## 📝 Notes

- NextAuth v5 beta (`5.0.0-beta.30`) menggunakan API yang berbeda dari v4
- Handler mengembalikan object dengan `GET` dan `POST` methods, bukan langsung function
- Response `null` dari session endpoint adalah normal ketika user belum login

