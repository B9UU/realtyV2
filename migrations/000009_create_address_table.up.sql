CREATE TABLE IF NOT EXISTS address (
  id bigserial PRIMARY KEY,
  listing_id integer NOT NULL REFERENCES listings(id) ON DELETE CASCADE,
  country text NOT NULL,
  province text NOT NULL,
  wijk text NOT NULL,
  neighbourhood text NOT NULL,
  house_number_suffix text NOT NULL,
  municipality text NOT NULL,
  is_bag_address boolean NOT NULL,
  house_number text NOT NULL,
  street_name text NOT NULL
);
