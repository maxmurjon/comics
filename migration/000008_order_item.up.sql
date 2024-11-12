CREATE TABLE order_items (
    order_id INT REFERENCES orders(id),
    comic_id INT REFERENCES comics(id),
    quantity INT NOT NULL DEFAULT 1,
    price DECIMAL(10, 2) NOT NULL,  -- Komiks narxi
    PRIMARY KEY (order_id, comic_id)
);
