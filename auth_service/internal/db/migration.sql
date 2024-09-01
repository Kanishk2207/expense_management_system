CREATE TABLE IF NOT EXISTS users (
    user_id VARCHAR(60) PRIMARY KEY,
    username VARCHAR(60) NOT NULL,
    first_name VARCHAR(60),
    last_name VARCHAR(60),
    email VARCHAR(60) NOT NULL,
    password VARCHAR(150) NOT NULL,
    created_at INT UNSIGNED DEFAULT 0,
    updated_at INT UNSIGNED DEFAULT 0
);

