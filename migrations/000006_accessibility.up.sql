CREATE TABLE IF NOT EXISTS accessibility (
  id bigserial PRIMARY KEY,
  text text NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS accessibilities (
  accessibility_id integer NOT NULL REFERENCES accessibility(id) ON DELETE CASCADE,
  listing_id integer NOT NULL REFERENCES listings(id) ON DELETE CASCADE,
  PRIMARY KEY(accessibility_id, listing_id)
);
