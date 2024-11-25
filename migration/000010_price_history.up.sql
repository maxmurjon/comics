CREATE TABLE price_history (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    old_price NUMERIC(10, 2) NOT NULL, -- Eski narx
    new_price NUMERIC(10, 2) NOT NULL, -- Yangi narx
    changed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- O'zgarish vaqti
);
