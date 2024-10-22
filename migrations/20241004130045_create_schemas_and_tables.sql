-- +goose Up
-- SQL запросы для создания схем и таблиц

-- Таблица authors в схеме library
CREATE TABLE IF NOT EXISTS authors
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-- Таблица books в схеме library
CREATE TABLE IF NOT EXISTS books
(
    id        SERIAL PRIMARY KEY,
    title     VARCHAR(200) NOT NULL,
    author_id INT REFERENCES authors (id) ON DELETE CASCADE
);

-- Таблица rentals_info в схеме library
CREATE TABLE IF NOT EXISTS rentals_info
(
    id          SERIAL PRIMARY KEY,
    user_id     INT,
    book_id     INT REFERENCES books (id) ON DELETE CASCADE,
    rental_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    return_date TIMESTAMP
);

-- Таблица rentals в схеме library
CREATE TABLE IF NOT EXISTS rentals
(
    id         INT PRIMARY KEY REFERENCES books (id) ON DELETE CASCADE,
    rentals_id INT REFERENCES rentals_info (id) DEFAULT NULL
);

-- +goose Down
-- SQL запросы для отката миграции

DROP TABLE IF EXISTS rentals;
DROP TABLE IF EXISTS rentals_info;
DROP TABLE IF EXISTS books;
DROP TABLE IF EXISTS authors;

DROP SCHEMA IF EXISTS library CASCADE;


