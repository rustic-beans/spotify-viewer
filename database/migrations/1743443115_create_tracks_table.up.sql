CREATE TABLE IF NOT EXISTS tracks (
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

CREATE TABLE IF NOT EXISTS artist_tracks (
  artist_id TEXT NOT NULL,
  track_id  TEXT NOT NULL,
  PRIMARY KEY (artist_id, track_id),
  FOREIGN KEY (artist_id) REFERENCES artists (id) ON DELETE CASCADE,
  FOREIGN KEY (track_id) REFERENCES tracks (id) ON DELETE CASCADE
);
