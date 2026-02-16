# Masjid Agung Discovery Residence - Landing Page Frontend

Landing page untuk Masjid Agung Discovery Residence menggunakan Next.js 15 dengan App Router, TypeScript, Tailwind CSS v4, dan shadcn/ui.

## 🚀 Tech Stack

- **Framework**: Next.js 15 (App Router)
- **Language**: TypeScript
- **Styling**: Tailwind CSS v4
- **UI Components**: shadcn/ui
- **Animations**: Framer Motion
- **Data Fetching**: TanStack Query (React Query)
- **HTTP Client**: Axios
- **Form Handling**: React Hook Form + Zod
- **Icons**: Lucide React
- **Counter Animation**: react-countup

## 📁 Project Structure

```
frontend/web/
├── app/                    # Next.js App Router pages
│   ├── layout.tsx          # Root layout dengan SEO meta tags
│   ├── page.tsx            # Landing page (home)
│   ├── donate/             # Donasi page
│   │   └── page.tsx
│   └── contact/            # Kontak page
│       └── page.tsx
├── components/
│   ├── ui/                 # shadcn/ui components
│   ├── sections/           # Page sections
│   │   ├── hero-section.tsx
│   │   ├── about-section.tsx
│   │   ├── donation-counter.tsx
│   │   ├── events-section.tsx
│   │   ├── gallery-section.tsx
│   │   └── contact-section.tsx
│   └── layouts/
│       └── footer.tsx
├── lib/
│   ├── api/                # API service layer
│   │   ├── client.ts       # Axios client setup
│   │   ├── donations.ts
│   │   ├── events.ts
│   │   ├── gallery.ts
│   │   ├── banners.ts
│   │   └── contact.ts
│   └── providers/
│       └── query-provider.tsx  # TanStack Query provider
└── public/                 # Static assets
```

## 🎯 Features

### Landing Page Sections

1. **Hero Section**
   - Background video atau banner image dari API
   - Title, subtitle, dan CTA button
   - Framer Motion fade-up animations
   - Scroll indicator animation

2. **About Section**
   - Informasi tentang masjid
   - Statistik (tahun berdiri, jamaah aktif)
   - Image dengan placeholder

3. **Donation Counter**
   - Fetch real-time data dari `/donations/summary`
   - Animated counter menggunakan react-countup
   - Breakdown per kategori
   - Auto-refresh setiap 1 menit
   - CTA button ke halaman donasi

4. **Events Section**
   - Fetch events dari `/events`
   - Card grid layout
   - Hover animations
   - Link ke detail event

5. **Gallery Section**
   - Fetch gallery items dari `/gallery`
   - Masonry grid layout
   - Image hover effects
   - Responsive grid (2-4 columns)

6. **Contact Form**
   - React Hook Form dengan Zod validation
   - Honeypot field untuk spam protection
   - Error handling
   - Success/error messages
   - Note: Backend endpoint `/contact` perlu diimplementasikan

7. **Footer**
   - Informasi kontak
   - Social media links
   - WhatsApp button
   - Quick links

## 🛠 Setup & Installation

### Prerequisites

- Node.js 18+ dan npm
- Backend API running di `http://localhost:8080`

### Installation

```bash
# Install dependencies
npm install

# Copy environment file
cp .env.example .env

# Edit .env file dan set API URL jika berbeda
NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1
```

### Development

```bash
# Run development server
npm run dev

# Open http://localhost:3000
```

### Build

```bash
# Build for production
npm run build

# Start production server
npm start
```

## 📡 API Integration

### Environment Variables

```env
NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1
```

### API Endpoints Used

- `GET /banners` - Hero banner/video
- `GET /donations/summary` - Donation summary (public)
- `GET /events` - List events
- `GET /gallery` - Gallery images
- `POST /contact` - Submit contact form (needs backend implementation)

### API Service Layer

Semua API calls menggunakan TanStack Query untuk:

- Caching
- Automatic refetching
- Loading states
- Error handling

Example usage:

```typescript
import { useQuery } from "@tanstack/react-query";
import { donationApi } from "@/lib/api/donations";

const { data, isLoading, error } = useQuery({
  queryKey: ["donation-summary"],
  queryFn: () => donationApi.getSummary(),
});
```

## 🎨 Styling

### Tailwind CSS v4

Project menggunakan Tailwind CSS v4 dengan konfigurasi default dari Next.js.

### shadcn/ui Components

Components yang sudah diinstall:

- Button
- Card
- Input
- Textarea
- Label

Untuk menambah components:

```bash
npx shadcn@latest add [component-name]
```

### Custom Styling

- Consistent spacing scale
- Soft transitions dan motion
- Responsive design (mobile-first)
- Dark mode ready (optional)

## 🎭 Animations

Menggunakan Framer Motion untuk:

- Page transitions
- Scroll-triggered animations
- Hover effects
- Loading states

Example:

```tsx
<motion.div
  initial={{ opacity: 0, y: 30 }}
  whileInView={{ opacity: 1, y: 0 }}
  viewport={{ once: true }}
  transition={{ duration: 0.8 }}
>
  Content
</motion.div>
```

## 📱 Responsive Design

- Mobile-first approach
- Breakpoints: sm (640px), md (768px), lg (1024px), xl (1280px)
- Grid layouts adapt sesuai screen size
- Touch-friendly buttons dan interactions

## 🔍 SEO

### Meta Tags

- Title, description, keywords
- OpenGraph tags untuk social sharing
- Twitter Card tags
- Robots meta

### Performance

- Image optimization dengan Next.js Image component
- Lazy loading untuk images
- Preload critical resources
- Code splitting otomatis

## 🚧 TODO / Next Steps

### Immediate

- [ ] Implement backend `/contact` endpoint
- [ ] Add placeholder images untuk About section
- [ ] Implement event detail page (`/events/[id]`)
- [ ] Add gallery lightbox/modal
- [ ] Add loading skeletons untuk better UX

### Future Enhancements

- [ ] Payment gateway integration untuk donasi
- [ ] Event registration/RSVP functionality
- [ ] Newsletter subscription
- [ ] Multi-language support (i18n)
- [ ] Dark mode toggle
- [ ] Advanced filtering untuk events & gallery
- [ ] Search functionality
- [ ] Analytics integration (Google Analytics, etc.)

## 📝 Notes

### Contact Form

Contact form saat ini mengirim ke endpoint `/contact` yang perlu diimplementasikan di backend. Sementara ini, form akan show error message jika endpoint belum tersedia.

### Donation Counter

Counter menggunakan `react-countup` untuk smooth number animation. Data di-refetch setiap 1 menit untuk update real-time.

### Image Optimization

Semua images menggunakan Next.js `Image` component dengan:

- Automatic optimization
- Lazy loading
- Responsive sizing
- Placeholder blur

### API Error Handling

TanStack Query menangani error secara otomatis. Untuk better UX, bisa tambahkan error boundaries dan fallback UI.

## 🤝 Contributing

1. Follow existing code structure
2. Use TypeScript untuk type safety
3. Follow Tailwind CSS conventions
4. Add proper error handling
5. Test responsive design
6. Update documentation jika perlu

## 📄 License

Copyright © 2024 Masjid Agung Discovery Residence. All rights reserved.
