CREATE TABLE IF NOT EXISTS plot_area_range (
    id bigserial PRIMARY KEY,
    gte integer NOT NULL ,
    lte integer NOT NULL
);

ALTER TABLE listings
ADD CONSTRAINT fk_plot_area_range
FOREIGN KEY (plot_area_range_id)
REFERENCES plot_area_range(id)
ON DELETE SET NULL; 

