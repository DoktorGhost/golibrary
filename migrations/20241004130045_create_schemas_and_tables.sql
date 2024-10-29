-- +goose Up
-- SQL запросы для создания схем и таблиц

-- Таблица rentals_info в схеме library
CREATE TABLE IF NOT EXISTS rentals_info
(
    id          SERIAL PRIMARY KEY,
    user_id     INT,
    book_id     INT,
    rental_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    return_date TIMESTAMP
);

-- Таблица rentals в схеме library
CREATE TABLE IF NOT EXISTS rentals
(
    id         INT,
    rentals_id INT REFERENCES rentals_info (id) DEFAULT NULL
);

-- +goose Down
-- SQL запросы для отката миграции

DROP TABLE IF EXISTS rentals;
DROP TABLE IF EXISTS rentals_info;



