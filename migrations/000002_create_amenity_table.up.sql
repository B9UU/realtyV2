CREATE TABLE IF NOT EXISTS amenity (
  id bigserial PRIMARY KEY,
  text text NOT NULL UNIQUE
);
