# Masjid CMS - Admin Dashboard

Professional Content Management System for Masjid Al-Madr built with Next.js 15, TypeScript, and modern web technologies.

## 🚀 Tech Stack

- **Framework**: Next.js 15 (App Router)
- **Language**: TypeScript
- **Styling**: Tailwind CSS v4
- **UI Components**: shadcn/ui
- **State Management**: TanStack React Query v5
- **HTTP Client**: Axios
- **Form Handling**: React Hook Form + Zod
- **Authentication**: NextAuth.js (JWT strategy)
- **Charts**: Recharts
- **Icons**: Lucide React
- **Notifications**: Sonner
- **Theme**: next-themes (Dark/Light mode)

## 📁 Project Structure

```
frontend/cms/
├── app/
│   ├── (auth)/
│   │   └── login/
│   │       └── page.tsx          # Login page
│   ├── api/
│   │   └── auth/
│   │       └── [...nextauth]/
│   │           └── route.ts     # NextAuth API route
│   ├── dashboard/
│   │   ├── layout.tsx           # Dashboard layout (sidebar + navbar)
│   │   ├── page.tsx              # Dashboard home (analytics)
│   │   ├── events/
│   │   │   └── page.tsx         # Events CRUD page
│   │   ├── gallery/
│   │   │   └── page.tsx         # Gallery CRUD page
│   │   ├── banner/
│   │   │   └── page.tsx         # Banner CRUD page
│   │   └── donations/
│   │       └── page.tsx         # Donations CRUD page
│   ├── layout.tsx                # Root layout
│   └── page.tsx                  # Root redirect
├── components/
│   ├── ui/                       # shadcn/ui components
│   ├── layout/
│   │   ├── sidebar.tsx          # Sidebar navigation
│   │   └── navbar.tsx           # Top navigation bar
│   ├── auth/
│   │   └── protected-route.tsx  # Route protection wrapper
│   ├── dashboard/
│   │   └── donations-chart.tsx  # Analytics chart component
│   └── events/
│       ├── events-table.tsx     # Events table component
│       └── event-modal.tsx      # Event create/edit modal
├── hooks/
│   ├── use-events.ts            # Events React Query hooks
│   └── use-stats.ts             # Dashboard stats hooks
├── lib/
│   ├── api/
│   │   ├── client.ts            # Axios client setup
│   │   ├── auth.ts              # Auth API
│   │   ├── events.ts            # Events API
│   │   ├── gallery.ts            # Gallery API
│   │   ├── banners.ts           # Banners API
│   │   ├── donations.ts         # Donations API
│   │   ├── stats.ts             # Stats API
│   │   ├── upload.ts            # File upload API
│   │   └── auth-interceptor.tsx # Auth token interceptor
│   ├── providers/
│   │   ├── query-provider.tsx   # TanStack Query provider
│   │   └── theme-provider.tsx   # Theme provider
│   ├── auth.ts                  # NextAuth configuration
│   └── logger.ts                # Centralized logging service
└── types/
    └── next-auth.d.ts           # NextAuth type definitions
```

## 🛠 Installation & Setup

### Prerequisites

- Node.js 18+ and npm
- Backend API running at `http://localhost:8080`

### Installation Steps

1. **Install dependencies:**

```bash
npm install
```

2. **Setup environment variables:**

```bash
cp .env.example .env
```

Edit `.env` file:

```env
# API Configuration
NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1

# NextAuth Configuration
NEXTAUTH_SECRET=your-secret-key-change-in-production
NEXTAUTH_URL=http://localhost:3001
```

**Important**: Change `NEXTAUTH_SECRET` to a secure random string in production. You can generate one using:

```bash
openssl rand -base64 32
```

3. **Run development server:**

```bash
npm run dev
```

The CMS will be available at `http://localhost:3001`

4. **Build for production:**

```bash
npm run build
npm start
```

## 🔐 Authentication

### Login Credentials

Default admin credentials (seeded by backend):

- **Username**: `admin`
- **Password**: `admin123`

### How Authentication Works

1. User submits login form with username/password
2. Frontend calls `/auth/login` endpoint
3. Backend returns JWT access token and refresh token
4. NextAuth stores tokens in JWT session
5. Access token is automatically added to all API requests via interceptor
6. Protected routes check authentication status

### Protected Routes

All routes under `/dashboard` are protected. Unauthenticated users are redirected to `/login`.

## 📡 API Integration

### API Base URL

Configured via `NEXT_PUBLIC_API_URL` environment variable.

Default: `http://localhost:8080/api/v1`

### API Endpoints Used

#### Authentication
- `POST /auth/login` - Login with credentials

#### Events
- `GET /admin/events` - List events (paginated)
- `GET /admin/events/:id` - Get event by ID
- `POST /admin/events` - Create event
- `PUT /admin/events/:id` - Update event
- `DELETE /admin/events/:id` - Delete event

#### Gallery
- `GET /admin/gallery` - List gallery items (paginated)
- `POST /admin/gallery` - Create gallery item (with file upload)
- `DELETE /admin/gallery/:id` - Delete gallery item

#### Banners
- `GET /admin/banners` - List banners (paginated)
- `GET /admin/banners/:id` - Get banner by ID
- `POST /admin/banners` - Create banner (with file upload)
- `PUT /admin/banners/:id` - Update banner
- `DELETE /admin/banners/:id` - Delete banner

#### Donations
- `GET /admin/donations` - List donations (paginated, filterable by status)
- `GET /admin/donations/:id` - Get donation by ID
- `POST /admin/donations` - Create donation
- `PUT /admin/donations/:id` - Update donation
- `DELETE /admin/donations/:id` - Delete donation
- `GET /admin/donation-categories` - List donation categories

#### File Upload
- `POST /admin/upload` - Upload file (image/video)

### API Client Setup

The API client (`lib/api/client.ts`) includes:

- Base URL configuration
- Request/response interceptors
- Automatic auth token injection
- Centralized error handling
- Request/response logging (development only)

### Using API Services

Example:

```typescript
import { eventApi } from "@/lib/api/events";

// In a component or hook
const { data, isLoading } = useQuery({
  queryKey: ["events"],
  queryFn: () => eventApi.getAll(10, 0),
});
```

## 🎨 Features

### Dashboard Home

- **Statistics Cards**: Total events, banners, gallery images, donations
- **Donations Chart**: Bar chart showing donations by category (using Recharts)

### CRUD Operations

Each module (Events, Gallery, Banner, Donations) includes:

- **List View**: Table with pagination, search, and filters
- **Create**: Modal form with validation
- **Edit**: Pre-filled modal form
- **Delete**: Confirmation dialog

### UI Features

- **Responsive Design**: Mobile-first, works on all screen sizes
- **Dark Mode**: System-based theme with manual toggle
- **Loading States**: Skeleton loaders and loading indicators
- **Error Handling**: Toast notifications for success/error
- **Form Validation**: Zod schema validation with React Hook Form
- **Optimistic Updates**: React Query mutations with cache invalidation

## 🔧 Development

### Adding New API Endpoints

1. Create API service in `lib/api/`:

```typescript
// lib/api/example.ts
import { apiClient, ApiResponse } from "./client";

export interface Example {
  id: number;
  name: string;
}

export const exampleApi = {
  getAll: async (): Promise<Example[]> => {
    const response = await apiClient.get<ApiResponse<Example[]>>("/admin/examples");
    return response.data.data || [];
  },
};
```

2. Create React Query hooks in `hooks/`:

```typescript
// hooks/use-examples.ts
import { useQuery } from "@tanstack/react-query";
import { exampleApi } from "@/lib/api/example";

export function useExamples() {
  return useQuery({
    queryKey: ["examples"],
    queryFn: () => exampleApi.getAll(),
  });
}
```

3. Use in components:

```typescript
const { data, isLoading } = useExamples();
```

### Adding New Pages

1. Create page file in `app/dashboard/your-module/page.tsx`
2. Add menu item to `components/layout/sidebar.tsx`
3. Follow existing CRUD pattern

### Styling

- Uses Tailwind CSS v4
- shadcn/ui components for consistent UI
- Custom styles in `app/globals.css`
- Theme variables for dark mode support

## 📝 Logging

Centralized logging service (`lib/logger.ts`):

- **Development**: All logs (debug, info, warn, error)
- **Production**: Only info, warn, error (no debug)

Usage:

```typescript
import { logger } from "@/lib/logger";

logger.debug("Debug message", data);
logger.info("Info message");
logger.warn("Warning message");
logger.error("Error message", error);
logger.api("GET", "/api/endpoint", requestData);
```

## 🚀 Deployment

### Environment Variables

Ensure these are set in production:

- `NEXT_PUBLIC_API_URL` - Backend API URL
- `NEXTAUTH_SECRET` - Secure random secret
- `NEXTAUTH_URL` - Frontend URL (e.g., `https://cms.yourdomain.com`)

### Build & Deploy

```bash
npm run build
npm start
```

Or deploy to platforms like Vercel, Netlify, etc.

## 🐛 Troubleshooting

### Authentication Issues

- Check `NEXTAUTH_SECRET` is set
- Verify `NEXTAUTH_URL` matches your domain
- Check browser console for errors
- Verify backend API is running and accessible

### API Connection Issues

- Verify `NEXT_PUBLIC_API_URL` is correct
- Check CORS settings on backend
- Verify backend is running
- Check network tab in browser DevTools

### Build Errors

- Clear `.next` folder: `rm -rf .next`
- Clear node_modules: `rm -rf node_modules && npm install`
- Check TypeScript errors: `npm run build`

## 📚 Additional Resources

- [Next.js Documentation](https://nextjs.org/docs)
- [TanStack Query Documentation](https://tanstack.com/query/latest)
- [shadcn/ui Components](https://ui.shadcn.com)
- [NextAuth.js Documentation](https://next-auth.js.org)
- [React Hook Form](https://react-hook-form.com)
- [Zod Documentation](https://zod.dev)

## 🤝 Contributing

1. Follow existing code structure
2. Use TypeScript for type safety
3. Follow React Query patterns for data fetching
4. Add proper error handling
5. Update documentation if needed

## 📄 License

Copyright © 2024 Masjid Al-Madr. All rights reserved.
