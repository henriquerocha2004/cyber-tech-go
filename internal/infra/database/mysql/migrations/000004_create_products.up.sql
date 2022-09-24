CREATE TABLE products (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    product_group BIGINT UNSIGNED NOT NULL,
    min_quantity INTEGER NOT NULL,
    max_quantity INTEGER NOT NULL,
    cost_value DECIMAL(10,2) NOT NULL DEFAULT 0.00,
    value DECIMAL(10,2) NOT NULL DEFAULT 0.00,
    PRIMARY KEY (id)
)
CHARACTER SET utf8mb4
COLLATE utf8mb4_general_ci
ENGINE = innodb;

ALTER TABLE products ADD CONSTRAINT fk_product_group FOREIGN KEY (product_group) REFERENCES product_categories (id);
