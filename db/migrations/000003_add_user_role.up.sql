ALTER TABLE users
ADD COLUMN IF NOT EXISTS role VARCHAR(50) CHECK (role IN ('admin', 'user')) NOT NULL DEFAULT 'user';