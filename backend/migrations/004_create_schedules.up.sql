CREATE TABLE schedules (
                           id SERIAL PRIMARY KEY,
                           event_name VARCHAR(255) NOT NULL,
                           event_date TIMESTAMP NOT NULL,
                           location VARCHAR(255),
                           chef_id INT REFERENCES chefs(id) ON DELETE CASCADE
);
