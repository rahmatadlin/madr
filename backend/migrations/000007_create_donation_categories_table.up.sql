-- Create donation_categories table
CREATE TABLE IF NOT EXISTS donation_categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Create indexes
CREATE INDEX IF NOT EXISTS idx_donation_categories_name ON donation_categories(name);
CREATE INDEX IF NOT EXISTS idx_donation_categories_deleted_at ON donation_categories(deleted_at);

