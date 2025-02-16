CREATE TABLE movies (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT UNIQUE NOT NULL,
    score DECIMAL(3,1) NOT NULL,
    picture TEXT NOT NULL,
    release_date DATE NOT NULL,
    synopsis TEXT NOT NULL,
    publisher_id UUID NOT NULL REFERENCES publisher(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);


---- create above / drop below ----

DROP TABLE IF EXISTS movies;
