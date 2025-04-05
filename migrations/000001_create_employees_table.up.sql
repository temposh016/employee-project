
CREATE TABLE employees (
                           id SERIAL PRIMARY KEY,
                           first_name VARCHAR(100) NOT NULL,
                           last_name VARCHAR(100) NOT NULL,
                           email VARCHAR(100) UNIQUE NOT NULL,
                           position VARCHAR(100),
                           hired_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO employees (first_name, last_name, email, position) VALUES
                                                                   ('Duman', 'Dudu', 'duman.doe@example.com', 'Manager'),
                                                                   ('Era', 'Erer', 'era.smith@example.com', 'Developer'),
                                                                   ('Aigen', 'Aiai', 'aigen.johnson@example.com', 'Designer');
