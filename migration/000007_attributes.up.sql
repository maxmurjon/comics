CREATE TABLE attributes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    data_type VARCHAR(50) NOT NULL, -- Masalan, 'string', 'integer', 'float'
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- SELECT products.id,products.name,products.description,products.price,products.stock_quantity,
--     attributes.name,product_attributes.value,attributes.data_type,
--     products.created_at,products.updated_at 
-- FROM products 
--     LEFT JOIN  product_attributes
--     ON products.id=product_attributes.product_id
--     LEFT JOIN attributes
--     ON product_attributes.attribute_id=attributes.id;
