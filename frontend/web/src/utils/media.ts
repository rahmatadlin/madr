const apiBase = (import.meta.env.VITE_API_URL || '').replace(/\/$/, '')
const apiBaseWithoutApi = apiBase.replace(/\/api\/v1\/?$/, '')
const defaultBase = apiBaseWithoutApi || 'http://localhost:8080'

export function resolveMediaUrl(url: string): string {
  return url.startsWith('http') ? url : `${defaultBase}/uploads/${url}`
}
