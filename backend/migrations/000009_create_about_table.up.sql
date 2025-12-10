-- Create about table
CREATE TABLE IF NOT EXISTS about (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    subtitle VARCHAR(255),
    description TEXT,
    additional_description TEXT,
    image_url VARCHAR(500),
    years_active INT NOT NULL DEFAULT 0,
    active_members INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_about_deleted_at ON about(deleted_at);

-- Seed initial content if table is empty
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM about) THEN
        INSERT INTO about (title, subtitle, description, additional_description, image_url, years_active, active_members)
        VALUES (
            'Tentang Masjid Al-Madr',
            'Pusat kegiatan keagamaan dan sosial',
            'Masjid Al-Madr adalah pusat kegiatan keagamaan dan sosial yang berkomitmen untuk membangun komunitas yang harmonis dan berkualitas.',
            'Dengan dukungan dari para jamaah dan donatur, kami terus mengembangkan fasilitas dan program yang dapat memberikan manfaat lebih luas bagi masyarakat.',
            '',
            15,
            500
        );
    END IF;
END$$;

