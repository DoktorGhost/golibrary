-- +goose Up
-- SQL запросы для создания схем и таблиц

CREATE SCHEMA IF NOT EXISTS library;
CREATE SCHEMA IF NOT EXISTS users;


-- Таблица authors в схеме library
CREATE TABLE IF NOT EXISTS library.authors
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-- Таблица books в схеме library
CREATE TABLE IF NOT EXISTS library.books
(
    id        SERIAL PRIMARY KEY,
    title     VARCHAR(200) NOT NULL,
    author_id INT REFERENCES library.authors (id) ON DELETE CASCADE
);

-- Таблица users в схеме users
CREATE TABLE IF NOT EXISTS users.users
(
    id            SERIAL PRIMARY KEY,
    username      VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    full_name     VARCHAR(200) NOT NULL
);


-- Таблица rentals_info в схеме library
CREATE TABLE IF NOT EXISTS library.rentals_info
(
    id          SERIAL PRIMARY KEY,
    user_id     INT REFERENCES users.users (id) ON DELETE CASCADE,
    book_id     INT REFERENCES library.books (id) ON DELETE CASCADE,
    rental_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    return_date TIMESTAMP
);

-- Таблица rentals в схеме library
CREATE TABLE IF NOT EXISTS library.rentals
(
    id         INT PRIMARY KEY REFERENCES library.books (id) ON DELETE CASCADE,
    rentals_id INT REFERENCES library.rentals_info (id) DEFAULT NULL
);

-- +goose Down
-- SQL запросы для отката миграции

DROP TABLE IF EXISTS library.rentals;
DROP TABLE IF EXISTS library.rentals_info;
DROP TABLE IF EXISTS users.users;
DROP TABLE IF EXISTS library.books;
DROP TABLE IF EXISTS library.authors;

DROP SCHEMA IF EXISTS library CASCADE;
DROP SCHEMA IF EXISTS users CASCADE;

