ALTER TABLE IF NOT EXISTS users ADD CONSTRAINT unique_username UNIQUE (username);
CREATE INDEX IF NOT EXISTS idx_username ON users(username);
