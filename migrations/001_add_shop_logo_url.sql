-- Add logo_url column to shops table
ALTER TABLE shops ADD COLUMN IF NOT EXISTS logo_url TEXT DEFAULT '';
