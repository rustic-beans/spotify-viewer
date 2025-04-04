CREATE TABLE IF NOT EXISTS artists (
  id            TEXT NOT NULL PRIMARY KEY,
  external_urls JSON NOT NULL,
  href          TEXT NOT NULL,
  name          TEXT NOT NULL,
  uri           TEXT NOT NULL,
  genres        JSON NOT NULL
);

CREATE TABLE IF NOT EXISTS artist_images (
  artist_id TEXT NOT NULL,
  image_url  TEXT NOT NULL,
  PRIMARY KEY (artist_id, image_url),
  FOREIGN KEY (artist_id) REFERENCES artists (id) ON DELETE CASCADE,
  FOREIGN KEY (image_url) REFERENCES images (url) ON DELETE CASCADE
);  

