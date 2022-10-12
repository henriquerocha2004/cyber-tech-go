SET sql_mode='NO_ZERO_DATE';

CREATE TABLE order_service (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    number VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    user_id BIGINT UNSIGNED NOT NULL,
    close_date DATETIME NOT NULL DEFAULT '0000-00-00 00:00:00',
    status_id BIGINT UNSIGNED NOT NULL,
    paid BOOLEAN NOT NULL DEFAULT false,
    created_at DATETIME,
    updated_at DATETIME,
    total DECIMAL(10,2) NOT NULL DEFAULT 0.00,
    PRIMARY KEY (id)
)
CHARACTER SET utf8mb4
COLLATE utf8mb4_general_ci
ENGINE = innodb;

CREATE TABLE order_items (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    product_id BIGINT UNSIGNED NOT NULL,
    order_id BIGINT UNSIGNED NOT NULL,
    type VARCHAR(255) NOT NULL,
    quantity INTEGER NOT NULL,
    value DECIMAL(10,2) NOT NULL,
    PRIMARY KEY(id)
)
CHARACTER SET utf8mb4
COLLATE utf8mb4_general_ci
ENGINE = innodb;

CREATE TABLE equipments (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    description VARCHAR(255) NOT NULL,
    defect VARCHAR(255) NOT NULL,
    observations VARCHAR(255),
    order_id BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY(id)
)
CHARACTER SET utf8mb4
COLLATE utf8mb4_general_ci
ENGINE = innodb;


CREATE TABLE order_payments (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    order_id BIGINT UNSIGNED NOT NULL,
    description VARCHAR(255) NOT NULL,
    total_value DECIMAL(10,2) NOT NULL DEFAULT 0.00,
    installments INTEGER NOT NULL DEFAULT 1,
    installment_value DECIMAL(10,2) NOT NULL DEFAULT 0.00,
    PRIMARY KEY (id)
)
CHARACTER SET utf8mb4
COLLATE utf8mb4_general_ci
ENGINE = innodb;

ALTER TABLE order_service ADD CONSTRAINT fk_status_order FOREIGN KEY (status_id) REFERENCES order_service_status (id);

ALTER TABLE order_items ADD CONSTRAINT fk_order_items_product FOREIGN KEY (product_id) REFERENCES products (id);

ALTER TABLE order_items ADD CONSTRAINT fk_order_items_order FOREIGN KEY (order_id) REFERENCES order_service (id);

ALTER TABLE equipments ADD CONSTRAINT fk_equipments_order FOREIGN KEY (order_id) REFERENCES order_service (id);

ALTER TABLE order_payments ADD CONSTRAINT fk_order_payments_order FOREIGN KEY (order_id) REFERENCES order_service (id);