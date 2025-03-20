CREATE TABLE subscriptions (
                               user_id INT REFERENCES users(id) ON DELETE CASCADE,
                               chef_id INT REFERENCES chefs(id) ON DELETE CASCADE,
                               PRIMARY KEY (user_id, chef_id)
);
