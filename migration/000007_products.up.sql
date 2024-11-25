CREATE TABLE products (
    id SERIAL PRIMARY KEY, -- Mahsulotning unikal identifikatori
    name VARCHAR(255) NOT NULL, -- Mahsulot nomi
    description TEXT, -- Mahsulot tavsifi
    price NUMERIC(10, 2) NOT NULL, -- Mahsulot narxi
    stock_quantity INT NOT NULL DEFAULT 0, -- Ombordagi mahsulot miqdori
    category_id INT, -- Kategoriyaning ID si (chet el kaliti)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Yaratilgan vaqti
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
