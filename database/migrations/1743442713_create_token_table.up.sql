CREATE TABLE IF NOT EXISTS token (
  id INTEGER NOT NULL PRIMARY KEY,
  access_token TEXT NOT NULL,
  token_type TEXT NOT NULL,
  refresh_token TEXT NOT NULL,
  expiry TIMESTAMP NOT NULL
);
