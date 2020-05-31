CREATE DATABASE database_test;

USE database_test;

CREATE TABLE IF NOT EXISTS freedom_fighters (
  id bigint NOT NULL AUTO_INCREMENT,
  first_name varchar(25),
  last_name varchar(25),
  work_at varchar(25),
  timestamp timestamp,
  PRIMARY KEY (id)
);

INSERT INTO freedom_fighters (id, first_name, last_name, work_at, timestamp)
VALUES
  (null, 'Richard', 'Stallman', 'FSF', now()),
  (null, 'Mitch', 'Kapor', 'EFF', now()),
  (null, 'John', 'Gilmore', 'EFF', now()),
  (null, 'John', 'Perry Barlow', 'EFF', now()),
  (null, 'Aaron', 'Swartz', 'CC', now());
