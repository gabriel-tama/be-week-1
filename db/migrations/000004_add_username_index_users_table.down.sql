

ALTER TABLE users DROP CONSTRAINT IF EXISTS unique_username;

DROP INDEX IF EXISTS username_idx;
