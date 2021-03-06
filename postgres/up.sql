DROP TABLE IF EXISTS gophers;

CREATE TABLE gophers (
  id VARCHAR(255) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  color VARCHAR(255) DEFAULT NULL,
  weight FLOAT DEFAULT NULL,
  image VARCHAR(255) DEFAULT NULL, -- Relational field with images table (id)
  created_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE images (
  id VARCHAR(255) PRIMARY KEY,
  content TEXT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL
);