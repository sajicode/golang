create table posts (
  id      serial primary key,
  content text,
  author  varchar(255)
);

-- to create table from the cmd line..
-- enter into the psql shell
-- psql -h localhost -U gwp
-- run `\ir setup.sql` to run sql script