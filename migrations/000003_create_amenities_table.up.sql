CREATE TABLE IF NOT EXISTS amenities (
  amenity_id integer,
  listing_id integer,
  CONSTRAINT fk_amenity FOREIGN KEY(amenity_id) REFERENCES amenity("id"),
  CONSTRAINT fk_listing FOREIGN KEY(listing_id) REFERENCES listing("id")
);
