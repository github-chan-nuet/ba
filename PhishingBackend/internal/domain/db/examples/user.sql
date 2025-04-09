-- SELECT pg_encoding_to_char(encoding) FROM pg_database WHERE datname = current_database();

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
insert into users values (uuid_generate_v4(), 'test', '\x4c26730b8d9e68fe64bf1d029d36f5842def8b79ace77f8c3e0d61daf700ce00'::bytea, 'infomail', 'simulationmail');