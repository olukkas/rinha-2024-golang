CREATE TABLE clients (
  id SERIAL PRIMARY KEY,
  balance INTEGER NOT NULL,
  total_limit INTEGER NOT NULL
);

CREATE TABLE transactions (
  id SERIAL PRIMARY KEY,
  client_id INTEGER REFERENCES clients(id),
  value INTEGER NOT NULL,
  type "char" CHECK (type IN ('c', 'd')) NOT NULL,
  description VARCHAR(10) NOT NULL,
  created_at TIMESTAMPTZ DEFAULT current_timestamp NOT NULL
);
