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
  relative_url text NOT NULL,
  plot_area_range_id integer NOT NULL
);
