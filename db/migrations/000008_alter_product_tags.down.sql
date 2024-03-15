DROP INDEX IF EXISTS tags_index;


ALTER TABLE product_tags
DROP COLUMN is_deleted;