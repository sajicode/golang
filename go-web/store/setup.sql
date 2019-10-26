drop table posts cascade if exists;
drop table comments if exists;

create table posts (
  id      serial primary key,
  content text,
  author  varchar(255)
);

create table comments (
  id      serial primary key,
  content text,
  author  varchar(255),
  post_id integer references posts(id)
);

-- to create table from the cmd line..
-- enter into the psql shell
-- psql -h localhost -U gwp
-- run `\ir setup.sql` to run sql script
-- when we drop the post tables, we need to cascade it, otherwise we won't be able to drop posts bcos the comments table depends on it