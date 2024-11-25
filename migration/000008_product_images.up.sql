CREATE TABLE product_images (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL REFERENCES products(id) ON DELETE CASCADE, -- Mahsulot ID si (chet el kaliti)
    image_url TEXT NOT NULL, -- Rasm URL manzili
    is_primary BOOLEAN DEFAULT FALSE -- Asosiy rasm belgisi
);
