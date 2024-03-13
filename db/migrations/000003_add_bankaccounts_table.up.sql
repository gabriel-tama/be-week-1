CREATE TABLE IF NOT EXISTS "bankaccounts"(
    account_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    bank_name VARCHAR(255) NOT NULL ,
    account_name VARCHAR(255) NOT NULL,
    account_number VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES Users(id)

);