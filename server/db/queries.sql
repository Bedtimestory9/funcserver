-- sudo -u lawrence psql postgres psql

-- create table
CREATE TABLE users (
username varchar(40) NOT NULL,
password varchar(40) NOT NULL
);

-- select table
SELECT * FROM user;

-- alter table
ALTER TABLE users
ADD age integer; 

-- insert to table
INSERT INTO users (username, password) VALUES 
    ('admin', 'admin1234'),
    ('admin2', 'admin1235');

-- delete
DELETE FROM golang_table WHERE name = 'Lawrence';

-- drop table
-- DROP TABLE user_table;
