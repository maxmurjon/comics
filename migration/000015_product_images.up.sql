CREATE TABLE product_images (
    id SERIAL PRIMARY KEY,                          -- Unikal identifikator
    product_id INT REFERENCES products(id) ON DELETE CASCADE, -- Mahsulot bilan bog'lanish
    image_url TEXT NOT NULL,                        -- Rasm URL manzili
    is_primary BOOLEAN DEFAULT FALSE,              -- Asosiy rasm flagi
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Yaratilgan vaqt
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- Yangilangan vaqt
);
