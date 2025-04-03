ALTER TABLE albums DROP COLUMN image_url;

CREATE TABLE IF NOT EXISTS album_images (
  album_id TEXT NOT NULL,
  image_url TEXT NOT NULL,
  PRIMARY KEY (album_id, image_url),
  FOREIGN KEY (album_id) REFERENCES albums (id) ON DELETE CASCADE,
  FOREIGN KEY (image_url) REFERENCES images (url) ON DELETE CASCADE
);
