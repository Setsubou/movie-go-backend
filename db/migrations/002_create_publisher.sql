CREATE TABLE publisher (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    publisher_name TEXT UNIQUE NOT NULL,
    year_founded int NOT NULL,
    country_id UUID NOT NULL REFERENCES country(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);


---- create above / drop below ----

DROP TABLE IF EXISTS publisher;
