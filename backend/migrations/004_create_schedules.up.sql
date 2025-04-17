CREATE TABLE schedules (
                           id SERIAL PRIMARY KEY,
                           event_name VARCHAR(255) NOT NULL, -- название маршрута
                           event_date INTEGER NOT NULL, -- оценка
                           location TEXT, -- текст отзыва
                           chef_id INT REFERENCES chefs(id) ON DELETE CASCADE
);
