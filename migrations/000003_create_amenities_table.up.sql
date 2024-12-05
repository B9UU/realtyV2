CREATE TABLE IF NOT EXISTS amenities (
  amenity_id integer NOT NULL REFERENCES amenity(id) ON DELETE CASCADE,
  listing_id integer NOT NULL REFERENCES listings(id) ON DELETE CASCADE,
  PRIMARY KEY(amenity_id, listing_id)
);

INSERT INTO amenities (amenity_id, listing_id) VALUES (1, 1);
INSERT INTO amenities (amenity_id, listing_id) VALUES (2, 1);
INSERT INTO amenities (amenity_id, listing_id) VALUES (3, 1);
