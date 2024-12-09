CREATE TABLE IF NOT EXISTS accessibility (
  id bigserial PRIMARY KEY,
  text text NOT NULL
);

CREATE TABLE IF NOT EXISTS accessibilities (
  accessibility_id integer NOT NULL REFERENCES amenity(id) ON DELETE CASCADE,
  listing_id integer NOT NULL REFERENCES listings(id) ON DELETE CASCADE,
  PRIMARY KEY(accessibility_id, listing_id)
);
