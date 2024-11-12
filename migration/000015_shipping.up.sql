CREATE TABLE shipping (
    id SERIAL PRIMARY KEY,
    order_id INT REFERENCES orders(id),
    shipping_address TEXT,
    shipping_date TIMESTAMPTZ,
    delivery_date TIMESTAMPTZ,
    status VARCHAR(50) DEFAULT 'pending',  -- Shipping holati (pending, delivered, canceled)
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);