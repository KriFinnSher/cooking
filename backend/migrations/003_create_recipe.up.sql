CREATE TABLE recipes (
                         id SERIAL PRIMARY KEY,
                         user_id INT REFERENCES users(id) ON DELETE CASCADE,
                         title VARCHAR(255) NOT NULL,
                         ingredients JSONB,
                         recipe_text TEXT
);
