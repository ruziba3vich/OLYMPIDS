CREATE TABLE athletes (
    id UUID PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    gender VARCHAR(50),
    nationality VARCHAR(255),
    height VARCHAR(50),
    weight VARCHAR(50),
    sport VARCHAR(255),
    date_of_birth TIMESTAMP,
    created_at TIMESTAMP DEFAULT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL
);
