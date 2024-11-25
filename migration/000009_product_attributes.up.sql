CREATE TABLE product_attributes (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    key VARCHAR(255) NOT NULL, -- Atribut nomi (masalan, "Rang")
    value VARCHAR(255) NOT NULL -- Atribut qiymati (masalan, "Qizil")
);
