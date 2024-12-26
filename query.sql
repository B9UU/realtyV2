-- name: Search :many
SELECT 
    sqlc.embed(listing)
    -- ARRAY_AGG(DISTINCT offering_type.text) FILTER (WHERE offering_type.text IS NOT NULL) as offering_type,
    -- ARRAY_AGG(DISTINCT mtp.text) FILTER (WHERE mtp.text IS NOT NULL) as media_types,
    -- ARRAY_AGG(DISTINCT amnt.text) FILTER (WHERE amnt.text IS NOT NULL) as amenities,
    -- ARRAY_AGG(DISTINCT acc.text) FILTER (WHERE acc.text IS NOT NULL) as accessibility,
    -- ARRAY_AGG(DISTINCT srnd.text) FILTER (WHERE srnd.text IS NOT NULL) as surrounding,
    -- ARRAY_AGG(DISTINCT parking.text) FILTER (WHERE parking.text IS NOT NULL) as parking_facility,
    --
    -- COALESCE(row_to_json(par),'{}')as plot_area_range,
    -- COALESCE(json_agg(DISTINCT agnt) FILTER (WHERE agnt.id IS NOT NULL), '[]') AS agents,
    -- COALESCE(row_to_json(address),'{}') as address

FROM listings l
LEFT JOIN amenity_listing amnts ON amnts.listing_id = l.id
LEFT JOIN amenity amnt ON amnts.amenity_id = amnt.id

LEFT JOIN media_type_listing mtps ON mtps.listing_id = l.id
LEFT JOIN media_type mtp ON mtps.media_type_id = mtp.id

LEFT JOIN agent_listing agnts ON agnts.listing_id = l.id
LEFT JOIN agent agnt ON agnts.agent_id = agnt.id

LEFT JOIN accessibility_listing accs ON accs.listing_id =l.id
LEFT JOIN accessibility acc ON acc.id = accs.accessibility_id

LEFT JOIN surrounding_listing srnds ON srnds.listing_id = l.id
LEFT JOIN surrounding srnd ON srnds.surrounding_id = srnd.id

LEFT JOIN address ON address.listing_id = l.id

LEFT JOIN plot_area_range par ON par.id = l.plot_area_range_id

LEFT JOIN parking_listing ON parking_listing.listing_id = l.id
LEFT JOIN parking ON parking_listing.parking_id = parking.id

LEFT JOIN offering_type_listing ON offering_type_listing.listing_id = l.id
LEFT JOIN offering_type ON offering_type.id = offering_type_listing.offering_type_id

WHERE geohash && ST_MakeEnvelope($1, $2, $3, $4, 4326)
GROUP BY l.id, par.id, address.id;

-- name: Ggs :many
SELECT id
FROM listings
WHERE id = $1;
