-- name: GetAlbums :many
SELECT *
FROM albums
ORDER BY name;

-- name: GetAlbumById :one
SELECT *
FROM albums
WHERE id = $1;

-- name: CreateAlbum :one
INSERT INTO albums (id, album_type, total_tracks, external_urls, href, name, release_date, release_date_precision, uri, genres)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: GetArtists :many
SELECT *
FROM artists
ORDER BY name;

-- name: GetArtistById :one
SELECT *
FROM artists
WHERE id = $1;

-- name: CreateArtist :one
INSERT INTO artists (id, external_urls, href, name, uri, genres)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetImages :many
SELECT *
FROM images
ORDER BY url;

-- name: GetImageByUrl :one
SELECT *
FROM images
WHERE url = $1;

-- name: CreateImage :one
INSERT INTO images (url, width, height)
VALUES ($1, $2, $3) ON CONFLICT DO NOTHING
RETURNING *;

-- name: GetTracks :many
SELECT *
FROM tracks
ORDER BY name;

-- name: GetTrackById :one
SELECT *
FROM tracks
WHERE id = $1;

-- name: CreateTrack :one
INSERT INTO tracks (id, duration_ms, explicit, external_urls, href, name, popularity, preview_url, track_number, uri, album_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING *;

-- name: SetAlbumImage :exec
INSERT INTO album_images (album_id, image_url)
VALUES ($1, $2);

-- name: SetArtistAlbum :exec
INSERT INTO artist_albums (artist_id, album_id)
VALUES ($1, $2);

-- name: SetArtistTrack :exec
INSERT INTO artist_tracks (artist_id, track_id)
VALUES ($1, $2);

-- name: SetArtistImage :exec
INSERT INTO artist_images (artist_id, image_url)
VALUES ($1, $2);
