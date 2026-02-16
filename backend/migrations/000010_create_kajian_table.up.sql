-- Kajian table for storing YouTube videos synced from channel
CREATE TABLE IF NOT EXISTS kajian (
    id SERIAL PRIMARY KEY,
    video_id VARCHAR(20) NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    published_at TIMESTAMP WITH TIME ZONE NOT NULL,
    thumbnail_url VARCHAR(512),
    youtube_url VARCHAR(512) NOT NULL,
    channel_title VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_kajian_deleted_at ON kajian(deleted_at);
CREATE INDEX IF NOT EXISTS idx_kajian_video_id ON kajian(video_id);
CREATE INDEX IF NOT EXISTS idx_kajian_published_at ON kajian(published_at DESC);
