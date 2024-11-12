CREATE TABLE comic_pages (
    id SERIAL PRIMARY KEY,
    comic_id INT REFERENCES comics(id),
    page_number INT NOT NULL,  -- Sahifa raqami
    page_url TEXT NOT NULL,  -- Sahifa manzili (pdf yoki boshqa formatda)
    created_at TIMESTAMPTZ DEFAULT NOW()
);