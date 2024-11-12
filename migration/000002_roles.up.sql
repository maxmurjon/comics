CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,  -- Masalan, 'user', 'admin', 'premium'
    description TEXT
);
