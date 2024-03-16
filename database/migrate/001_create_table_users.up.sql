CREATE TABLE users (
    id VARBINARY(1000) NOT NULL PRIMARY KEY,
    id_bit_count BIGINT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email_id VARCHAR(100) NOT NULL,
    country_code VARCHAR(10) NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT users_email_id_unq UNIQUE (email_id),
    CONSTRAINT users_country_code_phone_number UNIQUE (country_code, phone_number)
);