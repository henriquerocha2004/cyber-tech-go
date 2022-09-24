CREATE TABLE product_categories (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    description VARCHAR(255) NOT NULL,
    PRIMARY KEY(id)
) 
CHARACTER SET utf8mb4
COLLATE utf8mb4_general_ci
ENGINE = innodb;