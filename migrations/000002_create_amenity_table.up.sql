CREATE TABLE IF NOT EXISTS amenity (
  id bigserial PRIMARY KEY,
  amenity varchar,
  text varchar
);
INSERT INTO amenity (amenity, text) VALUES ('Pool','pool');
INSERT INTO amenity (amenity, text) VALUES ('Gym', 'gym');
INSERT INTO amenity (amenity, text) VALUES ('Parking','parking');
