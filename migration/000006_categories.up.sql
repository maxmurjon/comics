CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    parent_id INT, -- Agar sub-kategoriya bo'lsa, asosiy kategoriyaning ID si
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
