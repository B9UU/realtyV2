CREATE TABLE IF NOT EXISTS  parking(
    id bigserial PRIMARY KEY,
    text text NOT NULL
);

CREATE TABLE IF NOT EXISTS  parkings(
  listing_id int NOT NULL REFERENCES listings(id) ON DELETE CASCADE,
  parking_id int NOT NULL REFERENCES parking(id) ON DELETE CASCADE,
  PRIMARY KEY(parking_id, listing_id)
);
