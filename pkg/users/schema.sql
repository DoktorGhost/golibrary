CREATE TABLE IF NOT EXISTS users
(
    id            SERIAL PRIMARY KEY,
    username      VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    full_name     VARCHAR(200) NOT NULL
);
