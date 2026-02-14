const base = (import.meta.env.VITE_API_URL || '').replace(/\/$/, '').replace(/\/api\/v1\/?$/, '') || 'http://localhost:8080'

export function resolveMediaUrl(url?: string | null): string | null {
  if (!url) return null
  return url.startsWith('http') ? url : `${base}/uploads/${url}`
}
