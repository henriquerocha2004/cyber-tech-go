CREATE TABLE order_service_status (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    description VARCHAR(255) NOT NULL,
    launch_financial bool NOT NULL DEFAULT false,
    color VARCHAR(255),
    PRIMARY KEY (id)
)
CHARACTER SET utf8mb4
COLLATE utf8mb4_general_ci
ENGINE = innodb;