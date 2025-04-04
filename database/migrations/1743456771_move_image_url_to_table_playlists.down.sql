ALTER TABLE playlists DROP COLUMN image_url;

CREATE TABLE IF NOT EXISTS playlist_images (
  playlist_id TEXT NOT NULL,
  image_url  TEXT NOT NULL,
  PRIMARY KEY (playlist_id, image_url),
  FOREIGN KEY (playlist_id) REFERENCES playlists (id) ON DELETE CASCADE,
  FOREIGN KEY (image_url) REFERENCES images (url) ON DELETE CASCADE
);

