CREATE TABLE IF NOT EXISTS "payment" (
    id SERIAL PRIMARY KEY,
    account_id INT NOT NULL REFERENCES bankaccounts(account_id),
    product_id INT NOT NULL REFERENCES product(id),
    payment_proof TEXT NOT NULL,
    quantity INT NOT NULL CHECK (quantity >= 1)
);

CREATE INDEX IF NOT EXISTS idx_payment_account_id ON "payment" (account_id);
CREATE INDEX IF NOT EXISTS idx_payment_product_id ON "payment" (product_id);
