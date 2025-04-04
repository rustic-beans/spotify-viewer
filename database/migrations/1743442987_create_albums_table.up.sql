CREATE TYPE album_release_date_precision AS ENUM (
  'year',
  'month',
  'day'
);

CREATE TYPE album_type AS ENUM (
  'album',
  'single',
  'compilation'
);

CREATE TABLE IF NOT EXISTS albums (
  id                     TEXT NOT NULL PRIMARY KEY,
  album_type             album_type NOT NULL,
  total_tracks           BIGINT NOT NULL,
  external_urls          JSON NOT NULL,
  href                   TEXT NOT NULL,
  name                   TEXT NOT NULL,
  release_date           TEXT NOT NULL,
  release_date_precision album_release_date_precision NOT NULL,
  uri                    TEXT NOT NULL,
  genres                 JSON NOT NULL
);

CREATE TABLE IF NOT EXISTS album_images (
  album_id TEXT NOT NULL,
  image_url TEXT NOT NULL,
  PRIMARY KEY (album_id, image_url),
  FOREIGN KEY (album_id) REFERENCES albums (id) ON DELETE CASCADE,
  FOREIGN KEY (image_url) REFERENCES images (url) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS artist_albums (
  artist_id TEXT NOT NULL,
  album_id  TEXT NOT NULL,
  PRIMARY KEY (artist_id, album_id),
  FOREIGN KEY (artist_id) REFERENCES artists (id) ON DELETE CASCADE,
  FOREIGN KEY (album_id) REFERENCES albums (id) ON DELETE CASCADE
);
