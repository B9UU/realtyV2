CREATE TABLE IF NOT EXISTS  surrounding(
    id bigserial PRIMARY KEY,
    text text NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS  surroundings(
  listing_id int NOT NULL REFERENCES listings(id) ON DELETE CASCADE,
  surrounding_id int NOT NULL REFERENCES surrounding(id) ON DELETE CASCADE,
  PRIMARY KEY(surrounding_id, listing_id)
);
