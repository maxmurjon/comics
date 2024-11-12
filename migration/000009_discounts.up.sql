CREATE TABLE discounts (
    id SERIAL PRIMARY KEY,
    code VARCHAR(50) UNIQUE NOT NULL,  -- Masalan, 'DISCOUNT10'
    discount_percentage INT NOT NULL,  -- Masalan, 10% chegirma
    valid_until TIMESTAMP
);