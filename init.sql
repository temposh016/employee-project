-- Миграция для таблицы employees
CREATE TABLE employees (
                           id SERIAL PRIMARY KEY,
                           name VARCHAR(100),
                           position VARCHAR(100),
                           department VARCHAR(100),
                           salary FLOAT,
                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                           updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                           deleted_at TIMESTAMP
);
