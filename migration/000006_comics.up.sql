CREATE TABLE comics (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    description TEXT,
    genre VARCHAR(100),
    release_date TIMESTAMPTZ,
    popularity_score INT DEFAULT 0,  -- Mashhurlik reytingi
    poster_url TEXT,
    price DECIMAL(10, 2) NOT NULL,  -- Komiks narxi
    is_active BOOLEAN DEFAULT TRUE,  -- Komiksning faolligi
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);