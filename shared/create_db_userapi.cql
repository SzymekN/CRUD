-- Create a keyspace
CREATE KEYSPACE IF NOT EXISTS userapi WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : '1' };

-- Create a table
CREATE TABLE IF NOT EXISTS userapi.users (
    id int PRIMARY KEY,
    firstname text,
    lastname text,
    age int
    );

-- Insert some data
insert into userapi.users(id, firstname,lastname, age) values(1, 'Szymon', 'Nowak', 22);
insert into userapi.users(id, firstname,lastname, age) values(2, 'Jan', 'Kowalski', 31);
insert into userapi.users(id, firstname,lastname, age) values(3, 'Chuck', 'Norris', 18);
insert into userapi.users(id, firstname,lastname, age) values(4, 'Andrzej', 'Duda', 41);