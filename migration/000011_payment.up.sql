CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    -- purchase_id INT REFERENCES purchases(id),
    amount DECIMAL(10, 2),
    payment_method VARCHAR(50),  -- Masalan, 'credit_card', 'paypal'
    payment_date TIMESTAMP DEFAULT NOW(),
    status VARCHAR(50)  -- 'completed', 'pending', 'failed'
);
