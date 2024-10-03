-- Таблица для хранения авторов
CREATE TABLE IF NOT EXISTS authors (
                                       id SERIAL PRIMARY KEY,
                                       name VARCHAR(100) NOT NULL
);

-- Таблица для хранения книг
CREATE TABLE IF NOT EXISTS books (
                                     id SERIAL PRIMARY KEY,
                                     title VARCHAR(200) NOT NULL,
                                     author_id INT REFERENCES authors(id) ON DELETE CASCADE
);

-- Таблица для хранения пользователей
CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     username VARCHAR(100) NOT NULL UNIQUE,
                                     password_hash VARCHAR(255) NOT NULL,
                                     full_name VARCHAR(200) NOT NULL
);

-- Таблица для хранения информации об аренде книг
CREATE TABLE IF NOT EXISTS rentals_info (
                                       id SERIAL PRIMARY KEY,
                                       user_id INT REFERENCES users(id) ON DELETE CASCADE,
                                       book_id INT REFERENCES books(id) ON DELETE CASCADE,
                                       rental_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                       return_date TIMESTAMP
);

-- Таблица аренды книг
CREATE TABLE IF NOT EXISTS rentals (
                                            id INT PRIMARY KEY REFERENCES books(id) ON DELETE CASCADE,
                                            rentals_id INT REFERENCES rentals_info(id)
);