CREATE SCHEMA test

CREATE TABLE test.links (
  id SERIAL PRIMARY KEY,
  url varchar(255) NOT NULL, 
  name varchar(255) NOT NULL
);

INSERT INTO 
    test.links (url, name)
VALUES
    ('https://www.google.com','Google'),
    ('https://www.yahoo.com','Yahoo'),
    ('https://www.bing.com','Bing');