-- CREATE USER userapi WITH PASSWORD 'userapi' CREATEDB LOGIN;

-- CREATE DATABASE userapi
--     WITH 
--     OWNER = userapi
--     ENCODING = 'UTF8'
--     LC_COLLATE = 'en_US.UTF-8'
--     LC_CTYPE = 'en_US.UTF-8'
--     TABLESPACE = pg_default
--     CONNECTION LIMIT = -1;

\c userapi;
GRANT ALL PRIVILEGES ON ALL tables IN SCHEMA public TO userapi;
create table users(
    id int not null,
    firstname varchar(50),
    lastname varchar(50),
    age int,
    PRIMARY KEY(id)
);

insert into users(id,firstname,lastname, age) values(1, 'Szymon', 'Nowak', 22);
insert into users(id,firstname,lastname, age) values(2, 'Jan', 'Kowalski', 31);
insert into users(id,firstname,lastname, age) values(3, 'Chuck', 'Norris', 18);
insert into users(id,firstname,lastname, age) values(4, 'Andrzej', 'Duda', 41);