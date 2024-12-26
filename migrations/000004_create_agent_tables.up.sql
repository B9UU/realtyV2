CREATE TABLE IF NOT EXISTS agent (
  id bigserial PRIMARY KEY,
  logo_type text NOT NULL,
  relative_url text NOT NULL,
  is_primary bool NOT NULL,
  logo_id integer NOT NULL,
  name text NOT NULL,
  association text NOT NULL
);

CREATE TABLE IF NOT EXISTS agent_listing (
  listing_id integer NOT NULL REFERENCES listings(id) ON DELETE CASCADE,
  agent_id integer NOT NULL REFERENCES agent(id) ON DELETE CASCADE,
  PRIMARY KEY(listing_id, agent_id)
);

