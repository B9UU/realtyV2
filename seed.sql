INSERT INTO plot_area_range (gte, lte) VALUES
(0, 100),
(101, 200),
(201, 300);

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
  relative_url,
  plot_area_range_id
) 
VALUES
  ('For Sale', 2, 3, 5, 1, 'A+', 'Available', 'Apartment', 'Residential',
  CURRENT_TIMESTAMP, 'House', 'New', CURRENT_TIMESTAMP, CURRENT_DATE, '/listing-1',1),
 ('For Rent', 1, 1, 2, 5, 'B', 'Available', 'Studio', 'Commercial',
  CURRENT_TIMESTAMP, 'Office', 'Renovated', CURRENT_TIMESTAMP, CURRENT_DATE, '/listing-2',2),
 ('For Lease', 3, 4, 6, 2, 'A', 'Occupied', 'Townhouse', 'Mixed Use',
  CURRENT_TIMESTAMP, 'Villa', 'Old', CURRENT_TIMESTAMP, CURRENT_DATE, '/listing-3',3);

INSERT INTO amenity (text)
VALUES
('Pool'),
('Gym'),
('Parking');

INSERT INTO amenities (listing_id, amenity_id) 
VALUES
(1, 1),
(1, 2),
(1, 3),
(2, 1),
(2, 2),
(2, 3),
(3, 1),
(3, 2),
(3, 3);

INSERT INTO agent (id, logo_type, relative_url, is_primary, logo_id, name, association) 
VALUES 
(1, 'new', '/makelaar/24751-geijsel-makelaardij/', true, 159520467, 'Geijsel Makelaardij', 'NVM'),
(2, 'regular', '/makelaar/12345-smith-agency/', false, 123456789, 'Smith Agency', 'NVM'),
(3, 'premium', '/makelaar/67890-jones-realty/', true, 987654321, 'Jones Realty', 'VBO');

INSERT INTO agents (listing_id, agent_id)
VALUES
(1, 1),
(1, 2),
(2, 2),
(2, 3),
(3, 1),
(3, 3);


INSERT INTO accessibility (text)
VALUES
('laqo'),
('ground_floor'),
('single_storey');

INSERT INTO accessibilities (listing_id, accessibility_id) 
VALUES
(1, 1),
(1, 2),
(2, 1),
(2, 2),
(3, 1),
(3, 2);

INSERT INTO media_type (text)
VALUES
('floor_plan'),
('photo_360'),
('video');

INSERT INTO media_types (listing_id, media_type_id) 
VALUES
(1, 1),
(1, 2),
(2, 1),
(2, 2),
(3, 1),
(3, 2);

INSERT INTO surrounding (text)
VALUES
('on_quiet_road'),
('in_residential_district'),
('unobstructed_view');

INSERT INTO surroundings (listing_id, surrounding_id) 
VALUES
(1, 1),
(1, 2),
(2, 1),
(2, 2),
(3, 1),
(3, 2);

INSERT INTO address (
    listing_id, 
    country, 
    province, 
    wijk, 
    neighbourhood, 
    house_number_suffix, 
    municipality, 
    is_bag_address, 
    house_number, 
    street_name
) VALUES 
(1, 'Netherlands', 'North Holland', 'Wijk 01', 'Neighbourhood A', 'A1', 'Amsterdam', true, '123', 'Main Street'),
(2, 'Netherlands', 'South Holland', 'Wijk 02', 'Neighbourhood B', 'B2', 'Rotterdam', false, '456', 'Second Avenue'),
(3, 'Netherlands', 'Utrecht', 'Wijk 03', 'Neighbourhood C', 'C3', 'Utrecht', true, '789', 'Third Boulevard');

INSERT INTO parking (text)
VALUES
('paid_parking'),
('parking_permits'),
('parking_rental');

INSERT INTO parkings (listing_id, parking_id) 
VALUES
(1, 1),
(1, 2),
(2, 1),
(2, 2),
(3, 1),
(3, 2);
