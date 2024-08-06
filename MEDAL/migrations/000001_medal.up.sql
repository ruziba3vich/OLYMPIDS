CREATE TYPE medal_type AS ENUM('bronze', 'silver', 'gold');
CREATE TABLE medals (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    description VARCHAR,
    athlete_id UUID DEFAULT NULL,
    type medal_type,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);
SELECT * FROM medals
WHERE created_at BETWEEN '2024-08-07 00:00:00' AND '2024-08-07 23:59:59';
