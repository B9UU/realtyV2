CREATE TABLE IF NOT EXISTS  thumbnail(
    id bigserial PRIMARY KEY,
    listing_id int NOT NULL REFERENCES listings(id) ON DELETE CASCADE,
    img int NOT NULL UNIQUE
);
