-- Wishlist Table (Foydalanuvchilarning orzu ro'yxati)
CREATE TABLE wishlists (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE, -- Foydalanuvchi ID-si
    product_id INT REFERENCES products(id) ON DELETE CASCADE, -- Mahsulot ID-si
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Qachon qo'shilganligi
    UNIQUE(user_id, product_id) -- Foydalanuvchi faqat bitta mahsulotni bir martada orzu ro'yxatiga qo'shishi mumkin
);
