CREATE TABLE IF NOT EXISTS amenity_listing (
  amenity_id integer NOT NULL REFERENCES amenity(id) ON DELETE CASCADE,
  listing_id integer NOT NULL REFERENCES listings(id) ON DELETE CASCADE,
  PRIMARY KEY(amenity_id, listing_id)
);
