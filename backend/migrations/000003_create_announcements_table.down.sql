-- Drop announcements table
DROP INDEX IF EXISTS idx_announcements_deleted_at;
DROP INDEX IF EXISTS idx_announcements_is_published;
DROP TABLE IF EXISTS announcements;

