CREATE TABLE IF NOT EXISTS  media_type(
    id bigserial PRIMARY KEY,
    text text NOT NULL
);

CREATE TABLE IF NOT EXISTS  media_types(
  listing_id int NOT NULL REFERENCES listings(id) ON DELETE CASCADE,
  media_type_id int NOT NULL REFERENCES media_type(id) ON DELETE CASCADE,
  PRIMARY KEY(media_type_id, listing_id)
);


