-- Active: 1659521641949@@127.0.0.1@3306@cybertech

CREATE TABLE users (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    document VARCHAR(255),
    password VARCHAR(255) NOT NULL,
    type_person VARCHAR(255),
    last_login DATETIME,
    created_at DATETIME,
    updated_at DATETIME,
    created_by BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY (id)  
) 
CHARACTER SET utf8mb4
COLLATE utf8mb4_general_ci
ENGINE = innodb;

CREATE TABLE addresses (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    street VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    district VARCHAR(255) NOT NULL,
    state VARCHAR(255) NOT NULL,
    country VARCHAR(255) NOT NULL,
    zip_code VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    user_id BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY (id)
) 
CHARACTER SET utf8mb4
COLLATE utf8mb4_general_ci
ENGINE = innodb;

CREATE TABLE contacts (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    type VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    is_whatsapp BOOLEAN NOT NULL DEFAULT false,
    user_id BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY (id)
)
CHARACTER SET utf8mb4
COLLATE utf8mb4_general_ci
ENGINE = innodb; 

ALTER TABLE addresses ADD CONSTRAINT fk_address_user FOREIGN KEY (user_id) REFERENCES users (id);
ALTER TABLE contacts ADD CONSTRAINT fk_contacts_user FOREIGN KEY (user_id) REFERENCES users (id);