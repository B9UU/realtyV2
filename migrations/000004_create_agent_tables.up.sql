CREATE TABLE IF NOT EXISTS agent (
  id bigserial PRIMARY KEY,
  logo_type text,
  relative_url text,
  is_primary bool,
  logo_id integer,
  name text,
  association text
);

CREATE TABLE IF NOT EXISTS agents (
  listing_id integer NOT NULL REFERENCES listings(id) ON DELETE CASCADE,
  agent_id integer NOT NULL REFERENCES listings(id) ON DELETE CASCADE,
  PRIMARY KEY(listing_id, agent_id)
);

INSERT INTO agent (id, logo_type, relative_url, is_primary, logo_id, name, association) 
VALUES 
(1, 'new', '/makelaar/24751-geijsel-makelaardij/', true, 159520467, 'Geijsel Makelaardij', 'NVM');

INSERT INTO agent (id, logo_type, relative_url, is_primary, logo_id, name, association) 
VALUES 
(2, 'regular', '/makelaar/12345-smith-agency/', false, 123456789, 'Smith Agency', 'NVM');

INSERT INTO agent (id, logo_type, relative_url, is_primary, logo_id, name, association) 
VALUES 
(3, 'premium', '/makelaar/67890-jones-realty/', true, 987654321, 'Jones Realty', 'VBO');


-- Listing 1 is associated with Agent 1 and Agent 2
INSERT INTO agents (listing_id, agent_id) VALUES (1, 1);
INSERT INTO agents (listing_id, agent_id) VALUES (1, 2);

-- Listing 2 is associated with Agent 2 and Agent 3
INSERT INTO agents (listing_id, agent_id) VALUES (2, 2);
INSERT INTO agents (listing_id, agent_id) VALUES (2, 3);

-- Listing 3 is associated with Agent 1 and Agent 3
INSERT INTO agents (listing_id, agent_id) VALUES (3, 1);
INSERT INTO agents (listing_id, agent_id) VALUES (3, 3);
