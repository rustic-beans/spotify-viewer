package services

import (
	"context"
	"fmt"

	"github.com/rustic-beans/spotify-viewer/internal/database"
	"github.com/rustic-beans/spotify-viewer/internal/models"

	"github.com/jackc/pgx/v5"
)

type IDatabase interface {
	GetAlbums(ctx context.Context) ([]*models.Album, error)
	GetAlbumById(ctx context.Context, id string) (*models.Album, error)
	CreateAlbum(ctx context.Context, album *database.CreateAlbumParams, imageURLs []string, artistIDs []string) (*models.Album, error)

	GetArtists(ctx context.Context) ([]*models.Artist, error)
	GetArtistById(ctx context.Context, id string) (*models.Artist, error)
	CreateArtist(ctx context.Context, artist *database.CreateArtistParams, imageURLs []string) (*models.Artist, error)

	GetImages(ctx context.Context) ([]*models.Image, error)
	GetImageByUrl(ctx context.Context, url string) (*models.Image, error)
	CreateImages(ctx context.Context, image []*database.CreateImageParams) ([]*models.Image, error)

	GetTracks(ctx context.Context) ([]*models.Track, error)
	GetTrackById(ctx context.Context, id string) (*models.Track, error)
	CreateTrack(ctx context.Context, track *database.CreateTrackParams, artistIDs []string) (*models.Track, error)
}

type Database struct {
	*database.Queries
	client *pgx.Conn
}

func NewDatabase(client *pgx.Conn) IDatabase {
	return &Database{
		Queries: database.New(client),
		client:  client,
	}
}

func wrapOneQueryError[T any](result *T, err error) (*T, error) {
	if err == pgx.ErrNoRows {
		return nil, nil
	}

	return result, err
}

func wrapManyQueryError[T any](result []*T, err error) ([]*T, error) {
	if err == pgx.ErrNoRows {
		return nil, nil
	}

	return result, err
}

func (d *Database) withTX(ctx context.Context, fn func(*database.Queries) error) error {
	tx, err := d.client.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	qtx := d.Queries.WithTx(tx)
	if err := fn(qtx); err != nil {
		err := fmt.Errorf("error with transaction: %w", err)
		return err
	}

	return tx.Commit(ctx)
}

func (d *Database) GetAlbums(ctx context.Context) ([]*models.Album, error) {
	res, err := d.Queries.GetAlbums(ctx)
	return wrapManyQueryError(res, err)
}

func (d *Database) GetAlbumById(ctx context.Context, id string) (*models.Album, error) {
	res, err := d.Queries.GetAlbumById(ctx, id)
	return wrapOneQueryError(res, err)
}

func (d *Database) CreateAlbum(ctx context.Context, album *database.CreateAlbumParams, imageURLs []string, artistIDs []string) (*models.Album, error) {
	var a *models.Album
	err := d.withTX(ctx, func(q *database.Queries) error {
		var err error
		a, err = q.CreateAlbum(ctx, album)
		if err != nil {
			return err
		}

		for _, url := range imageURLs {
			err = q.SetAlbumImage(ctx, &database.SetAlbumImageParams{
				AlbumID:  a.ID,
				ImageUrl: url,
			})
			if err != nil {
				return err
			}
		}

		for _, id := range artistIDs {
			err = q.SetArtistAlbum(ctx, &database.SetArtistAlbumParams{
				AlbumID:  a.ID,
				ArtistID: id,
			})
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (d *Database) GetArtists(ctx context.Context) ([]*models.Artist, error) {
	res, err := d.Queries.GetArtists(ctx)
	return wrapManyQueryError(res, err)
}

func (d *Database) GetArtistById(ctx context.Context, id string) (*models.Artist, error) {
	res, err := d.Queries.GetArtistById(ctx, id)
	return wrapOneQueryError(res, err)
}

func (d *Database) CreateArtist(ctx context.Context, artist *database.CreateArtistParams, imageURLs []string) (*models.Artist, error) {
	var a *models.Artist
	err := d.withTX(ctx, func(q *database.Queries) error {
		var err error
		a, err = q.CreateArtist(ctx, artist)
		if err != nil {
			return err
		}

		for _, url := range imageURLs {
			err = q.SetArtistImage(ctx, &database.SetArtistImageParams{
				ArtistID: a.ID,
				ImageUrl: url,
			})
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (d *Database) GetImages(ctx context.Context) ([]*models.Image, error) {
	res, err := d.Queries.GetImages(ctx)
	return wrapManyQueryError(res, err)
}

func (d *Database) GetImageByUrl(ctx context.Context, url string) (*models.Image, error) {
	res, err := d.Queries.GetImageByUrl(ctx, url)
	return wrapOneQueryError(res, err)
}

func (d *Database) CreateImages(ctx context.Context, images []*database.CreateImageParams) ([]*models.Image, error) {
	var imgs []*models.Image
	err := d.withTX(ctx, func(q *database.Queries) error {
		imgs = make([]*models.Image, 0, len(images))
		for _, img := range images {
			i, err := q.CreateImage(ctx, img)
			if err != nil {
				return err
			}

			imgs = append(imgs, i)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return imgs, nil
}

func (d *Database) GetTracks(ctx context.Context) ([]*models.Track, error) {
	res, err := d.Queries.GetTracks(ctx)
	return wrapManyQueryError(res, err)
}

func (d *Database) GetTrackById(ctx context.Context, id string) (*models.Track, error) {
	res, err := d.Queries.GetTrackById(ctx, id)
	return wrapOneQueryError(res, err)
}

func (d *Database) CreateTrack(ctx context.Context, track *database.CreateTrackParams, artistIDs []string) (*models.Track, error) {
	var t *models.Track
	err := d.withTX(ctx, func(q *database.Queries) error {
		var err error
		t, err = q.CreateTrack(ctx, track)
		if err != nil {
			return err
		}

		for _, id := range artistIDs {
			err = q.SetArtistTrack(ctx, &database.SetArtistTrackParams{
				ArtistID: id,
				TrackID:  t.ID,
			})
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return t, nil
}
