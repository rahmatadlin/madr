-- Fix kajian table: remove invalid data and ensure video_id is NOT NULL
-- Delete any rows where video_id is NULL (invalid data from old schema)
DELETE FROM kajian WHERE video_id IS NULL;

-- Ensure video_id column is NOT NULL (in case it was created as nullable)
ALTER TABLE kajian ALTER COLUMN video_id SET NOT NULL;

-- Ensure unique constraint exists
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_constraint WHERE conname = 'kajian_video_id_key'
    ) THEN
        ALTER TABLE kajian ADD CONSTRAINT kajian_video_id_key UNIQUE (video_id);
    END IF;
END$$;
