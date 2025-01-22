// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package database

import (
	"context"
)

const createAlbum = `-- name: CreateAlbum :one
INSERT INTO albums (id, album_type, total_tracks, external_urls, href, name, release_date, release_date_precision, uri, genres)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING id, album_type, total_tracks, external_urls, href, name, release_date, release_date_precision, uri, genres
`

type CreateAlbumParams struct {
	ID                   string                    `json:"id"`
	AlbumType            AlbumType                 `json:"album_type"`
	TotalTracks          int64                     `json:"total_tracks"`
	ExternalUrls         map[string]string         `json:"external_urls"`
	Href                 string                    `json:"href"`
	Name                 string                    `json:"name"`
	ReleaseDate          string                    `json:"release_date"`
	ReleaseDatePrecision AlbumReleaseDatePrecision `json:"release_date_precision"`
	Uri                  string                    `json:"uri"`
	Genres               []string                  `json:"genres"`
}

func (q *Queries) CreateAlbum(ctx context.Context, arg *CreateAlbumParams) (*Album, error) {
	row := q.db.QueryRow(ctx, createAlbum,
		arg.ID,
		arg.AlbumType,
		arg.TotalTracks,
		arg.ExternalUrls,
		arg.Href,
		arg.Name,
		arg.ReleaseDate,
		arg.ReleaseDatePrecision,
		arg.Uri,
		arg.Genres,
	)
	var i Album
	err := row.Scan(
		&i.ID,
		&i.AlbumType,
		&i.TotalTracks,
		&i.ExternalUrls,
		&i.Href,
		&i.Name,
		&i.ReleaseDate,
		&i.ReleaseDatePrecision,
		&i.Uri,
		&i.Genres,
	)
	return &i, err
}

const createArtist = `-- name: CreateArtist :one
INSERT INTO artists (id, external_urls, href, name, uri, genres)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, external_urls, href, name, uri, genres
`

type CreateArtistParams struct {
	ID           string            `json:"id"`
	ExternalUrls map[string]string `json:"external_urls"`
	Href         string            `json:"href"`
	Name         string            `json:"name"`
	Uri          string            `json:"uri"`
	Genres       []string          `json:"genres"`
}

func (q *Queries) CreateArtist(ctx context.Context, arg *CreateArtistParams) (*Artist, error) {
	row := q.db.QueryRow(ctx, createArtist,
		arg.ID,
		arg.ExternalUrls,
		arg.Href,
		arg.Name,
		arg.Uri,
		arg.Genres,
	)
	var i Artist
	err := row.Scan(
		&i.ID,
		&i.ExternalUrls,
		&i.Href,
		&i.Name,
		&i.Uri,
		&i.Genres,
	)
	return &i, err
}

const createImage = `-- name: CreateImage :one
INSERT INTO images (url, width, height)
VALUES ($1, $2, $3) ON CONFLICT DO NOTHING
RETURNING url, width, height
`

type CreateImageParams struct {
	Url    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}

func (q *Queries) CreateImage(ctx context.Context, arg *CreateImageParams) (*Image, error) {
	row := q.db.QueryRow(ctx, createImage, arg.Url, arg.Width, arg.Height)
	var i Image
	err := row.Scan(&i.Url, &i.Width, &i.Height)
	return &i, err
}

const createTrack = `-- name: CreateTrack :one
INSERT INTO tracks (id, duration_ms, explicit, external_urls, href, name, popularity, preview_url, track_number, uri, album_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING id, duration_ms, explicit, external_urls, href, name, popularity, preview_url, track_number, uri, album_id
`

type CreateTrackParams struct {
	ID           string            `json:"id"`
	DurationMs   int64             `json:"duration_ms"`
	Explicit     bool              `json:"explicit"`
	ExternalUrls map[string]string `json:"external_urls"`
	Href         string            `json:"href"`
	Name         string            `json:"name"`
	Popularity   int64             `json:"popularity"`
	PreviewUrl   *string           `json:"preview_url"`
	TrackNumber  int64             `json:"track_number"`
	Uri          string            `json:"uri"`
	AlbumID      string            `json:"album_id"`
}

func (q *Queries) CreateTrack(ctx context.Context, arg *CreateTrackParams) (*Track, error) {
	row := q.db.QueryRow(ctx, createTrack,
		arg.ID,
		arg.DurationMs,
		arg.Explicit,
		arg.ExternalUrls,
		arg.Href,
		arg.Name,
		arg.Popularity,
		arg.PreviewUrl,
		arg.TrackNumber,
		arg.Uri,
		arg.AlbumID,
	)
	var i Track
	err := row.Scan(
		&i.ID,
		&i.DurationMs,
		&i.Explicit,
		&i.ExternalUrls,
		&i.Href,
		&i.Name,
		&i.Popularity,
		&i.PreviewUrl,
		&i.TrackNumber,
		&i.Uri,
		&i.AlbumID,
	)
	return &i, err
}

const getAlbumArtists = `-- name: GetAlbumArtists :many
SELECT artists.id, artists.external_urls, artists.href, artists.name, artists.uri, artists.genres
FROM artists
JOIN artist_albums ON artists.id = artist_albums.artist_id
WHERE artist_albums.album_id = $1
`

func (q *Queries) GetAlbumArtists(ctx context.Context, albumID string) ([]*Artist, error) {
	rows, err := q.db.Query(ctx, getAlbumArtists, albumID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Artist
	for rows.Next() {
		var i Artist
		if err := rows.Scan(
			&i.ID,
			&i.ExternalUrls,
			&i.Href,
			&i.Name,
			&i.Uri,
			&i.Genres,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAlbumImages = `-- name: GetAlbumImages :many
SELECT images.url, images.width, images.height
FROM images
JOIN album_images ON images.url = album_images.image_url
WHERE album_images.album_id = $1
`

func (q *Queries) GetAlbumImages(ctx context.Context, albumID string) ([]*Image, error) {
	rows, err := q.db.Query(ctx, getAlbumImages, albumID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Image
	for rows.Next() {
		var i Image
		if err := rows.Scan(&i.Url, &i.Width, &i.Height); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAlbumTracks = `-- name: GetAlbumTracks :many
SELECT id, duration_ms, explicit, external_urls, href, name, popularity, preview_url, track_number, uri, album_id
FROM tracks
WHERE album_id = $1
`

func (q *Queries) GetAlbumTracks(ctx context.Context, albumID string) ([]*Track, error) {
	rows, err := q.db.Query(ctx, getAlbumTracks, albumID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Track
	for rows.Next() {
		var i Track
		if err := rows.Scan(
			&i.ID,
			&i.DurationMs,
			&i.Explicit,
			&i.ExternalUrls,
			&i.Href,
			&i.Name,
			&i.Popularity,
			&i.PreviewUrl,
			&i.TrackNumber,
			&i.Uri,
			&i.AlbumID,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAlbums = `-- name: GetAlbums :many
SELECT id, album_type, total_tracks, external_urls, href, name, release_date, release_date_precision, uri, genres
FROM albums
ORDER BY name
`

func (q *Queries) GetAlbums(ctx context.Context) ([]*Album, error) {
	rows, err := q.db.Query(ctx, getAlbums)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Album
	for rows.Next() {
		var i Album
		if err := rows.Scan(
			&i.ID,
			&i.AlbumType,
			&i.TotalTracks,
			&i.ExternalUrls,
			&i.Href,
			&i.Name,
			&i.ReleaseDate,
			&i.ReleaseDatePrecision,
			&i.Uri,
			&i.Genres,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAlbumsById = `-- name: GetAlbumsById :many
SELECT id, album_type, total_tracks, external_urls, href, name, release_date, release_date_precision, uri, genres
FROM albums
WHERE id = ANY($1::text[])
`

func (q *Queries) GetAlbumsById(ctx context.Context, dollar_1 []string) ([]*Album, error) {
	rows, err := q.db.Query(ctx, getAlbumsById, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Album
	for rows.Next() {
		var i Album
		if err := rows.Scan(
			&i.ID,
			&i.AlbumType,
			&i.TotalTracks,
			&i.ExternalUrls,
			&i.Href,
			&i.Name,
			&i.ReleaseDate,
			&i.ReleaseDatePrecision,
			&i.Uri,
			&i.Genres,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getArtistAlbums = `-- name: GetArtistAlbums :many
SELECT albums.id, albums.album_type, albums.total_tracks, albums.external_urls, albums.href, albums.name, albums.release_date, albums.release_date_precision, albums.uri, albums.genres
FROM albums
JOIN artist_albums ON albums.id = artist_albums.album_id
WHERE artist_albums.artist_id = $1
`

func (q *Queries) GetArtistAlbums(ctx context.Context, artistID string) ([]*Album, error) {
	rows, err := q.db.Query(ctx, getArtistAlbums, artistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Album
	for rows.Next() {
		var i Album
		if err := rows.Scan(
			&i.ID,
			&i.AlbumType,
			&i.TotalTracks,
			&i.ExternalUrls,
			&i.Href,
			&i.Name,
			&i.ReleaseDate,
			&i.ReleaseDatePrecision,
			&i.Uri,
			&i.Genres,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getArtistImages = `-- name: GetArtistImages :many
SELECT images.url, images.width, images.height
FROM images
JOIN artist_images ON images.url = artist_images.image_url
WHERE artist_images.artist_id = $1
`

func (q *Queries) GetArtistImages(ctx context.Context, artistID string) ([]*Image, error) {
	rows, err := q.db.Query(ctx, getArtistImages, artistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Image
	for rows.Next() {
		var i Image
		if err := rows.Scan(&i.Url, &i.Width, &i.Height); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getArtistTracks = `-- name: GetArtistTracks :many
SELECT tracks.id, tracks.duration_ms, tracks.explicit, tracks.external_urls, tracks.href, tracks.name, tracks.popularity, tracks.preview_url, tracks.track_number, tracks.uri, tracks.album_id
FROM tracks
JOIN artist_tracks ON tracks.id = artist_tracks.track_id
WHERE artist_tracks.artist_id = $1
`

func (q *Queries) GetArtistTracks(ctx context.Context, artistID string) ([]*Track, error) {
	rows, err := q.db.Query(ctx, getArtistTracks, artistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Track
	for rows.Next() {
		var i Track
		if err := rows.Scan(
			&i.ID,
			&i.DurationMs,
			&i.Explicit,
			&i.ExternalUrls,
			&i.Href,
			&i.Name,
			&i.Popularity,
			&i.PreviewUrl,
			&i.TrackNumber,
			&i.Uri,
			&i.AlbumID,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getArtists = `-- name: GetArtists :many
SELECT id, external_urls, href, name, uri, genres
FROM artists
ORDER BY name
`

func (q *Queries) GetArtists(ctx context.Context) ([]*Artist, error) {
	rows, err := q.db.Query(ctx, getArtists)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Artist
	for rows.Next() {
		var i Artist
		if err := rows.Scan(
			&i.ID,
			&i.ExternalUrls,
			&i.Href,
			&i.Name,
			&i.Uri,
			&i.Genres,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getArtistsById = `-- name: GetArtistsById :many
SELECT id, external_urls, href, name, uri, genres
FROM artists
WHERE id = ANY($1::text[])
`

func (q *Queries) GetArtistsById(ctx context.Context, dollar_1 []string) ([]*Artist, error) {
	rows, err := q.db.Query(ctx, getArtistsById, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Artist
	for rows.Next() {
		var i Artist
		if err := rows.Scan(
			&i.ID,
			&i.ExternalUrls,
			&i.Href,
			&i.Name,
			&i.Uri,
			&i.Genres,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getImages = `-- name: GetImages :many
SELECT url, width, height
FROM images
ORDER BY url
`

func (q *Queries) GetImages(ctx context.Context) ([]*Image, error) {
	rows, err := q.db.Query(ctx, getImages)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Image
	for rows.Next() {
		var i Image
		if err := rows.Scan(&i.Url, &i.Width, &i.Height); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getImagesByUrl = `-- name: GetImagesByUrl :many
SELECT url, width, height
FROM images
WHERE url = ANY($1::text[])
`

func (q *Queries) GetImagesByUrl(ctx context.Context, dollar_1 []string) ([]*Image, error) {
	rows, err := q.db.Query(ctx, getImagesByUrl, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Image
	for rows.Next() {
		var i Image
		if err := rows.Scan(&i.Url, &i.Width, &i.Height); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTrackAlbum = `-- name: GetTrackAlbum :one
SELECT albums.id, albums.album_type, albums.total_tracks, albums.external_urls, albums.href, albums.name, albums.release_date, albums.release_date_precision, albums.uri, albums.genres
FROM albums
JOIN tracks ON albums.id = tracks.album_id
WHERE tracks.id = $1
`

func (q *Queries) GetTrackAlbum(ctx context.Context, id string) (*Album, error) {
	row := q.db.QueryRow(ctx, getTrackAlbum, id)
	var i Album
	err := row.Scan(
		&i.ID,
		&i.AlbumType,
		&i.TotalTracks,
		&i.ExternalUrls,
		&i.Href,
		&i.Name,
		&i.ReleaseDate,
		&i.ReleaseDatePrecision,
		&i.Uri,
		&i.Genres,
	)
	return &i, err
}

const getTrackArtists = `-- name: GetTrackArtists :many
SELECT artists.id, artists.external_urls, artists.href, artists.name, artists.uri, artists.genres
FROM artists
JOIN artist_tracks ON artists.id = artist_tracks.artist_id
WHERE artist_tracks.track_id = $1
`

func (q *Queries) GetTrackArtists(ctx context.Context, trackID string) ([]*Artist, error) {
	rows, err := q.db.Query(ctx, getTrackArtists, trackID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Artist
	for rows.Next() {
		var i Artist
		if err := rows.Scan(
			&i.ID,
			&i.ExternalUrls,
			&i.Href,
			&i.Name,
			&i.Uri,
			&i.Genres,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTracks = `-- name: GetTracks :many
SELECT id, duration_ms, explicit, external_urls, href, name, popularity, preview_url, track_number, uri, album_id
FROM tracks
ORDER BY name
`

func (q *Queries) GetTracks(ctx context.Context) ([]*Track, error) {
	rows, err := q.db.Query(ctx, getTracks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Track
	for rows.Next() {
		var i Track
		if err := rows.Scan(
			&i.ID,
			&i.DurationMs,
			&i.Explicit,
			&i.ExternalUrls,
			&i.Href,
			&i.Name,
			&i.Popularity,
			&i.PreviewUrl,
			&i.TrackNumber,
			&i.Uri,
			&i.AlbumID,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTracksById = `-- name: GetTracksById :many
SELECT id, duration_ms, explicit, external_urls, href, name, popularity, preview_url, track_number, uri, album_id
FROM tracks
WHERE id = ANY($1::text[])
`

func (q *Queries) GetTracksById(ctx context.Context, dollar_1 []string) ([]*Track, error) {
	rows, err := q.db.Query(ctx, getTracksById, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Track
	for rows.Next() {
		var i Track
		if err := rows.Scan(
			&i.ID,
			&i.DurationMs,
			&i.Explicit,
			&i.ExternalUrls,
			&i.Href,
			&i.Name,
			&i.Popularity,
			&i.PreviewUrl,
			&i.TrackNumber,
			&i.Uri,
			&i.AlbumID,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const setAlbumImage = `-- name: SetAlbumImage :exec
INSERT INTO album_images (album_id, image_url)
VALUES ($1, $2)
`

type SetAlbumImageParams struct {
	AlbumID  string `json:"album_id"`
	ImageUrl string `json:"image_url"`
}

func (q *Queries) SetAlbumImage(ctx context.Context, arg *SetAlbumImageParams) error {
	_, err := q.db.Exec(ctx, setAlbumImage, arg.AlbumID, arg.ImageUrl)
	return err
}

const setArtistAlbum = `-- name: SetArtistAlbum :exec
INSERT INTO artist_albums (artist_id, album_id)
VALUES ($1, $2)
`

type SetArtistAlbumParams struct {
	ArtistID string `json:"artist_id"`
	AlbumID  string `json:"album_id"`
}

func (q *Queries) SetArtistAlbum(ctx context.Context, arg *SetArtistAlbumParams) error {
	_, err := q.db.Exec(ctx, setArtistAlbum, arg.ArtistID, arg.AlbumID)
	return err
}

const setArtistImage = `-- name: SetArtistImage :exec
INSERT INTO artist_images (artist_id, image_url)
VALUES ($1, $2)
`

type SetArtistImageParams struct {
	ArtistID string `json:"artist_id"`
	ImageUrl string `json:"image_url"`
}

func (q *Queries) SetArtistImage(ctx context.Context, arg *SetArtistImageParams) error {
	_, err := q.db.Exec(ctx, setArtistImage, arg.ArtistID, arg.ImageUrl)
	return err
}

const setArtistTrack = `-- name: SetArtistTrack :exec
INSERT INTO artist_tracks (artist_id, track_id)
VALUES ($1, $2)
`

type SetArtistTrackParams struct {
	ArtistID string `json:"artist_id"`
	TrackID  string `json:"track_id"`
}

func (q *Queries) SetArtistTrack(ctx context.Context, arg *SetArtistTrackParams) error {
	_, err := q.db.Exec(ctx, setArtistTrack, arg.ArtistID, arg.TrackID)
	return err
}
