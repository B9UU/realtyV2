CREATE TABLE IF NOT EXISTS "amenities" (
  amenity_id integer,
  listing_id integer,
  FOREIGN KEY ("amenity_id") REFERENCES "amenity" ("id"),
  FOREIGN KEY ("listing_id") REFERENCES "listing" ("id")
);
