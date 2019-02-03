// psql -h localhost -p 5432 -U postgres
// \c go_playground

CREATE DATABASE go_playground;

CREATE TABLE projects (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  name VARCHAR (50) DEFAULT NULL,
  data TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO projects (user_id, name, data) VALUES
  (1, 'test house 1', 'data 1'),
  (1, 'test house 2', 'data 2'),
  (1, 'test house 3', 'data 3'),
  (1, 'test house 4', 'data 4'),
  (1, 'test house 5', 'data 5');
