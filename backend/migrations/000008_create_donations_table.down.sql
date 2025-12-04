-- Drop donations table
DROP INDEX IF EXISTS idx_donations_deleted_at;
DROP INDEX IF EXISTS idx_donations_payment_status;
DROP INDEX IF EXISTS idx_donations_category_id;
DROP TABLE IF EXISTS donations;

