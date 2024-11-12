CREATE TABLE permissions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,  -- Masalan, 'view_comic', 'buy_comic', 'edit_comic'
    description TEXT
);
