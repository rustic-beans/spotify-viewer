CREATE TABLE IF NOT EXISTS playlists (
  id           TEXT NOT NULL PRIMARY KEY,
  external_urls JSON NOT NULL,
  href          TEXT NOT NULL,
  name          TEXT NOT NULL,
  uri           TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS playlist_images (
  playlist_id TEXT NOT NULL,
  image_url  TEXT NOT NULL,
  PRIMARY KEY (playlist_id, image_url),
  FOREIGN KEY (playlist_id) REFERENCES playlists (id) ON DELETE CASCADE,
  FOREIGN KEY (image_url) REFERENCES images (url) ON DELETE CASCADE
);

