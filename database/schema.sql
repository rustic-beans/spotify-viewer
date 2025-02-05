CREATE TYPE album_type AS ENUM (
  'album',
  'single',
  'compilation'
);

CREATE TYPE album_release_date_precision AS ENUM (
  'year',
  'month',
  'day'
);

CREATE TABLE albums (
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

CREATE TABLE artists (
  id            TEXT NOT NULL PRIMARY KEY,
  external_urls JSON NOT NULL,
  href          TEXT NOT NULL,
  name          TEXT NOT NULL,
  uri           TEXT NOT NULL,
  genres        JSON NOT NULL
);

CREATE TABLE playlists (
  id           TEXT NOT NULL PRIMARY KEY,
  external_urls JSON NOT NULL,
  href          TEXT NOT NULL,
  name          TEXT NOT NULL,
  uri           TEXT NOT NULL
);

CREATE TABLE images (
  url    TEXT NOT NULL PRIMARY KEY,
  width  BIGINT NOT NULL,
  height BIGINT NOT NULL
);

CREATE TABLE tracks (
  id            TEXT NOT NULL PRIMARY KEY,
  duration_ms   BIGINT NOT NULL,
  explicit      BOOL NOT NULL DEFAULT (false),
  external_urls JSON NOT NULL,
  href          TEXT NOT NULL,
  name          TEXT NOT NULL,
  popularity    BIGINT NOT NULL,
  preview_url   TEXT NULL,
  track_number  BIGINT NOT NULL,
  uri           TEXT NOT NULL,
  album_id      TEXT NOT NULL,
  FOREIGN KEY (album_id) REFERENCES albums (id) ON DELETE no action
);

CREATE TABLE album_images (
  album_id TEXT NOT NULL,
  image_url TEXT NOT NULL,
  PRIMARY KEY (album_id, image_url),
  FOREIGN KEY (album_id) REFERENCES albums (id) ON DELETE CASCADE,
  FOREIGN KEY (image_url) REFERENCES images (url) ON DELETE CASCADE
);

CREATE TABLE artist_albums (
  artist_id TEXT NOT NULL,
  album_id  TEXT NOT NULL,
  PRIMARY KEY (artist_id, album_id),
  FOREIGN KEY (artist_id) REFERENCES artists (id) ON DELETE CASCADE,
  FOREIGN KEY (album_id) REFERENCES albums (id) ON DELETE CASCADE
);

CREATE TABLE artist_tracks (
  artist_id TEXT NOT NULL,
  track_id  TEXT NOT NULL,
  PRIMARY KEY (artist_id, track_id),
  FOREIGN KEY (artist_id) REFERENCES artists (id) ON DELETE CASCADE,
  FOREIGN KEY (track_id) REFERENCES tracks (id) ON DELETE CASCADE
);

CREATE TABLE artist_images
(
  artist_id TEXT NOT NULL,
  image_url  TEXT NOT NULL,
  PRIMARY KEY (artist_id, image_url),
  FOREIGN KEY (artist_id) REFERENCES artists (id) ON DELETE CASCADE,
  FOREIGN KEY (image_url) REFERENCES images (url) ON DELETE CASCADE
);  

CREATE TABLE playlist_images (
  playlist_id TEXT NOT NULL,
  image_url  TEXT NOT NULL,
  PRIMARY KEY (playlist_id, image_url),
  FOREIGN KEY (playlist_id) REFERENCES playlists (id) ON DELETE CASCADE,
  FOREIGN KEY (image_url) REFERENCES images (url) ON DELETE CASCADE
);
