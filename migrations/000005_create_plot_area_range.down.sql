ALTER TABLE listings
DROP CONSTRAINT IF EXISTS fk_plot_area_range;
ALTER TABLE listings
DROP CONSTRAINT IF EXISTS fk_floor_area_range;

DROP TABLE IF EXISTS plot_area_range;
DROP TABLE IF EXISTS floor_area_range;

