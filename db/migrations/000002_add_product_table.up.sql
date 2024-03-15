CREATE TABLE IF NOT EXISTS product (
    id SERIAL PRIMARY KEY,
    name VARCHAR(60) NOT NULL CHECK (LENGTH(name) >= 5),
    price DECIMAL NOT NULL CHECK (price >= 0),
    imageUrl TEXT NOT NULL,
    stock INT NOT NULL CHECK (stock >= 0),
    condition VARCHAR(10) NOT NULL CHECK (condition IN ('new', 'second')),
    isPurchaseable BOOLEAN NOT NULL
);

DROP TYPE IF EXISTS tag;
CREATE TYPE tag AS ENUM ();

CREATE TABLE IF NOT EXISTS product_tags (
    product_id INT REFERENCES product(id),
    tag tag NOT NULL
);