CREATE TYPE medal_type AS ENUM('bronze', 'silver', 'gold');
CREATE TABLE country_medals(
    name VARCHAR PRIMARY KEY UNIQUE NOT NULL,
    gold_count INT DEFAULT 0,
    silver_count INT DEFAULT 0,
    bronze_count INT DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);
CREATE TABLE medals (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    description VARCHAR,
    athlete_id UUID DEFAULT NULL,
    country VARCHAR ,
    type medal_type,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);
