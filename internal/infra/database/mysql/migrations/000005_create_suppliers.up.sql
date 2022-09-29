CREATE TABLE suppliers (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    document VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    district VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    state VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
)
ENGINE = innodb
COLLATE utf8mb4_general_ci
ENGINE = innodb;
