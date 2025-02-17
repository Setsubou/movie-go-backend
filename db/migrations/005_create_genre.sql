CREATE TABLE genres (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    genre TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);


---- create above / drop below ----

DROP TABLE IF EXISTS users;
