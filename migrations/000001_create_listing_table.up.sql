CREATE TABLE IF NOT EXISTS listings (
  id bigserial PRIMARY KEY,
  placement_type text NOT NULL,
  number_of_bathrooms integer NOT NULL,
  number_of_bedrooms integer NOT NULL,
  number_of_rooms integer NOT NULL,
  relevancy_sort_order integer NOT NULL,
  energy_label text NOT NULL,
  availability text NOT NULL,
  type text NOT NULL,
  zoning text NOT NULL,
  time_stamp timestamp NOT NULL,
  object_type text NOT NULL,
  construction_type text NOT NULL,
  publish_date_utc timestamp NOT NULL,
  publish_date timestamp NOT NULL,
  relative_url text
);

INSERT INTO listings (
  placement_type,
  number_of_bathrooms,
  number_of_bedrooms,
  number_of_rooms,
  relevancy_sort_order,
  energy_label,
  availability,
  type,
  zoning,
  time_stamp,
  object_type,
  construction_type,
  publish_date_utc,
  publish_date,
  relative_url
) VALUES (
  'For Sale', 2, 3, 5, 1, 'A+', 'Available', 'Apartment', 'Residential', CURRENT_TIMESTAMP, 'House', 'New', CURRENT_TIMESTAMP, CURRENT_DATE, '/listing-1'
);

INSERT INTO listings (
  placement_type,
  number_of_bathrooms,
  number_of_bedrooms,
  number_of_rooms,
  relevancy_sort_order,
  energy_label,
  availability,
  type,
  zoning,
  time_stamp,
  object_type,
  construction_type,
  publish_date_utc,
  publish_date,
  relative_url
) VALUES (
  'For Rent', 1, 1, 2, 5, 'B', 'Available', 'Studio', 'Commercial', CURRENT_TIMESTAMP, 'Office', 'Renovated', CURRENT_TIMESTAMP, CURRENT_DATE, '/listing-2'
);

INSERT INTO listings (
  placement_type,
  number_of_bathrooms,
  number_of_bedrooms,
  number_of_rooms,
  relevancy_sort_order,
  energy_label,
  availability,
  type,
  zoning,
  time_stamp,
  object_type,
  construction_type,
  publish_date_utc,
  publish_date,
  relative_url
) VALUES (
  'For Lease', 3, 4, 6, 2, 'A', 'Occupied', 'Townhouse', 'Mixed Use', CURRENT_TIMESTAMP, 'Villa', 'Old', CURRENT_TIMESTAMP, CURRENT_DATE, '/listing-3'
);
