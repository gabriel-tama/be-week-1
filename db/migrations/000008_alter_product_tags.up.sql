
CREATE INDEX IF NOT EXISTS tags_index  ON product_tags(tag);

ALTER TABLE product_tags ADD COLUMN is_deleted BOOLEAN NOT NULL DEFAULT false;