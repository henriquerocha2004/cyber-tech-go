CREATE TABLE stock (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    type_movement VARCHAR(255) NOT NULL,
    quantity INTEGER NOT NULL,
    invoice VARCHAR(255),
    date DATETIME NOT NULL,
    supplier_id BIGINT UNSIGNED NOT NULL,
    product_id BIGINT UNSIGNED NOT NULL,
    user_id BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY (id)
)
ENGINE = innodb
COLLATE utf8mb4_general_ci
ENGINE = innodb;

ALTER TABLE stock ADD CONSTRAINT fk_stock_supplier FOREIGN KEY (supplier_id) REFERENCES suppliers (id);
ALTER TABLE stock ADD CONSTRAINT fk_stock_product FOREIGN KEY (product_id) REFERENCES products (id);