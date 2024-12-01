CREATE TABLE languages (
    id SERIAL PRIMARY KEY,
    code VARCHAR(10) NOT NULL UNIQUE, -- Masalan, 'en', 'uz'
    name VARCHAR(255) NOT NULL -- Masalan, 'English', 'O'zbek'
);