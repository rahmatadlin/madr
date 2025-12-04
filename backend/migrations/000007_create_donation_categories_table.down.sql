-- Drop donation_categories table
DROP INDEX IF EXISTS idx_donation_categories_deleted_at;
DROP INDEX IF EXISTS idx_donation_categories_name;
DROP TABLE IF EXISTS donation_categories;

