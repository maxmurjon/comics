CREATE TABLE comic_reviews (
    id SERIAL PRIMARY KEY,
    comic_id INT REFERENCES comics(id),
    user_id INT REFERENCES users(id),
    rating INT CHECK (rating >= 1 AND rating <= 5),  -- 1 dan 5 gacha reyting
    review TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

