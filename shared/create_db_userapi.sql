-- CREATE USER userapi WITH PASSWORD 'userapi' CREATEDB LOGIN;

-- CREATE DATABASE userapi
--     WITH 
--     OWNER = userapi
--     ENCODING = 'UTF8'
--     LC_COLLATE = 'en_US.UTF-8'
--     LC_CTYPE = 'en_US.UTF-8'
--     TABLESPACE = pg_default
--     CONNECTION LIMIT = -1;

-- GRANT ALL PRIVILEGES ON ALL tables IN SCHEMA public TO userapi;
create table users(
    id serial,
    firstname varchar(50),
    lastname varchar(50),
    age int,
    PRIMARY KEY(id)
);

insert into users(firstname,lastname, age) values('Szymon', 'Nowak', 22);
insert into users(firstname,lastname, age) values('Jan', 'Kowalski', 31);
insert into users(firstname,lastname, age) values('Chuck', 'Norris', 18);
insert into users(firstname,lastname, age) values('Andrzej', 'Duda', 41);