CREATE TABLE IF NOT EXISTS offering_type(
    id bigserial PRIMARY KEY,
    text text NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS  offering_types(
  listing_id int NOT NULL REFERENCES listings(id) ON DELETE CASCADE,
  offering_type_id int NOT NULL REFERENCES offering_type(id) ON DELETE CASCADE,
  PRIMARY KEY(offering_type_id, listing_id)
);
