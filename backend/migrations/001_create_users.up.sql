CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(100) NOT NULL,
                       hash TEXT NOT NULL,
                       avatar TEXT
);
