CREATE TABLE chefs (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(100) NOT NULL,
                       hash TEXT NOT NULL,
                       followers_count INT DEFAULT 0,
                       bio TEXT,
                       avatar TEXT
);
