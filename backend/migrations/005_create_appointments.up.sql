CREATE TABLE appointments (
                              user_id INT REFERENCES users(id) ON DELETE CASCADE,
                              schedule_id INT REFERENCES schedules(id) ON DELETE CASCADE,
                              PRIMARY KEY (user_id, schedule_id)
);
