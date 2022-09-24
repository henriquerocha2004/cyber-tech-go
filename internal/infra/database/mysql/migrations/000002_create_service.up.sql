CREATE TABLE services (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    description VARCHAR(255),
    price DECIMAL(10,2),
    PRIMARY KEY (id)
)
CHARACTER SET utf8mb4
COLLATE utf8mb4_general_ci
ENGINE = innodb;