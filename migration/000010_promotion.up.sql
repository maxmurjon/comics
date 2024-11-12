CREATE TABLE promotions (
    id SERIAL PRIMARY KEY,
    comic_id INT REFERENCES comics(id),
    promotion_description TEXT,
    promotion_price DECIMAL(10, 2),
    valid_until TIMESTAMP
);
