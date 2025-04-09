-- select used string encoding (UTF-8, UTF-16 etc.)
 SELECT pg_encoding_to_char(encoding) FROM pg_database WHERE datname = current_database();

-- select prepared (https://stackoverflow.com/a/12160161)
SELECT relfilenode FROM pg_class WHERE relname = 'pg_prepared_statements';

-- use UUID extension so that function uuid_generate_v4 can be used
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- create test user with pw and user name "test"
insert into users values (uuid_generate_v4(), 'test', '\x4c26730b8d9e68fe64bf1d029d36f5842def8b79ace77f8c3e0d61daf700ce00'::bytea, 'infomail', 'simulationmail');