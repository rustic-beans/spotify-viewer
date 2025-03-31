-- name: GetAlbums :many
SELECT *
FROM albums
ORDER BY name;

-- name: GetAlbumsByID :many
SELECT *
FROM albums
WHERE id = ANY($1::text[]);

-- name: GetAlbumArtists :many
SELECT artists.*
FROM artists
JOIN artist_albums ON artists.id = artist_albums.artist_id
WHERE artist_albums.album_id = $1;

-- name: GetAlbumTracks :many
SELECT *
FROM tracks
WHERE album_id = $1;

-- name: CreateAlbum :one
INSERT INTO albums (id, album_type, total_tracks, external_urls, href, name, release_date, release_date_precision, uri, genres, image_url)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING *;

-- name: GetArtists :many
SELECT *
FROM artists
ORDER BY name;

-- name: GetArtistsByID :many
SELECT *
FROM artists
WHERE id = ANY($1::text[]);

-- name: GetArtistAlbums :many
SELECT albums.*
FROM albums
JOIN artist_albums ON albums.id = artist_albums.album_id
WHERE artist_albums.artist_id = $1;

-- name: GetArtistTracks :many
SELECT tracks.*
FROM tracks
JOIN artist_tracks ON tracks.id = artist_tracks.track_id
WHERE artist_tracks.artist_id = $1;

-- name: CreateArtist :one
INSERT INTO artists (id, external_urls, href, name, uri, genres, image_url)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetTracks :many
SELECT *
FROM tracks
ORDER BY name;

-- name: GetTracksByID :many
SELECT *
FROM tracks
WHERE id = ANY($1::text[]);

-- name: GetTrackAlbum :one
SELECT albums.*
FROM albums
JOIN tracks ON albums.id = tracks.album_id
WHERE tracks.id = $1;

-- name: GetTrackArtists :many
SELECT artists.*
FROM artists
JOIN artist_tracks ON artists.id = artist_tracks.artist_id
WHERE artist_tracks.track_id = $1;

-- name: CreateTrack :one
INSERT INTO tracks (id, duration_ms, explicit, external_urls, href, name, popularity, preview_url, track_number, uri, album_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING *;

-- name: SetArtistAlbum :exec
INSERT INTO artist_albums (artist_id, album_id)
VALUES ($1, $2);

-- name: SetArtistTrack :exec
INSERT INTO artist_tracks (artist_id, track_id)
VALUES ($1, $2);

-- name: GetPlaylists :many
SELECT *
FROM playlists
ORDER BY name;

-- name: GetPlaylistsByID :many
SELECT *
FROM playlists
WHERE id = ANY($1::text[]);

-- name: CreatePlaylist :one
INSERT INTO playlists (id, external_urls, href, name, uri, image_url)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpsertToken :one
INSERT INTO token (id, access_token, token_type, expiry, refresh_token)
VALUES (1, $1, $2, $3, $4)
ON CONFLICT (id) DO UPDATE
SET access_token = $1, token_type = $2, expiry = $3, refresh_token = $4
RETURNING *;

-- name: GetToken :one
SELECT *
FROM token
WHERE id = 1;
