ALTER TABLE listings
DROP CONSTRAINT IF EXISTS fk_plot_area_range;

DROP TABLE IF EXISTS property_area_range;
DROP TABLE IF EXISTS plot_area_range;

