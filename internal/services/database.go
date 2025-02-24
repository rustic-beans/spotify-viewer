package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/rustic-beans/spotify-viewer/internal/database"
	"github.com/rustic-beans/spotify-viewer/internal/models"
	"github.com/rustic-beans/spotify-viewer/internal/utils"
	"go.uber.org/zap"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IDatabase interface {
	HealthCheck(ctx context.Context) error

	GetAlbums(ctx context.Context) ([]*models.Album, error)
	GetAlbumsByID(ctx context.Context, id []string) ([]*models.Album, error)
	GetAlbumArtists(ctx context.Context, id string) ([]*models.Artist, error)
	GetAlbumImages(ctx context.Context, id string) ([]*models.Image, error)
	GetAlbumTracks(ctx context.Context, id string) ([]*models.Track, error)
	CreateAlbum(ctx context.Context, album *database.CreateAlbumParams, imageURLs []string, artistIDs []string) (*models.Album, error)

	GetArtists(ctx context.Context) ([]*models.Artist, error)
	GetArtistsByID(ctx context.Context, id []string) ([]*models.Artist, error)
	GetArtistAlbums(ctx context.Context, id string) ([]*models.Album, error)
	GetArtistImages(ctx context.Context, id string) ([]*models.Image, error)
	GetArtistTracks(ctx context.Context, id string) ([]*models.Track, error)
	CreateArtist(ctx context.Context, artist *database.CreateArtistParams, imageURLs []string) (*models.Artist, error)

	GetImages(ctx context.Context) ([]*models.Image, error)
	GetImagesByURL(ctx context.Context, url []string) ([]*models.Image, error)
	CreateImages(ctx context.Context, image []*database.CreateImageParams) ([]*models.Image, error)

	GetTracks(ctx context.Context) ([]*models.Track, error)
	GetTracksByID(ctx context.Context, id []string) ([]*models.Track, error)
	GetTrackAlbum(ctx context.Context, id string) (*models.Album, error)
	GetTrackArtists(ctx context.Context, id string) ([]*models.Artist, error)
	CreateTrack(ctx context.Context, track *database.CreateTrackParams, artistIDs []string) (*models.Track, error)

	GetPlaylists(ctx context.Context) ([]*models.Playlist, error)
	GetPlaylistsByID(ctx context.Context, id []string) ([]*models.Playlist, error)
	GetPlaylistImages(ctx context.Context, id string) ([]*models.Image, error)
	CreatePlaylist(ctx context.Context, playlist *database.CreatePlaylistParams, imageURLs []string) (*models.Playlist, error)

	UpsertToken(ctx context.Context, token *database.UpsertTokenParams) (*models.Token, error)
	GetToken(ctx context.Context) (*models.Token, error)
}

type Database struct {
	*database.Queries
	client *pgxpool.Pool
}

func NewDatabase(client *pgxpool.Pool) IDatabase {
	return &Database{
		Queries: database.New(client),
		client:  client,
	}
}

func (d *Database) HealthCheck(ctx context.Context) error {
	return d.client.Ping(ctx)
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

	defer func() {
		err = tx.Rollback(ctx)
		if err != nil {
			if errors.Is(err, pgx.ErrTxClosed) {
				return
			}

			utils.Logger.Error("error rolling back transaction", zap.Error(err))
		}
	}()

	qtx := d.Queries.WithTx(tx)
	if err = fn(qtx); err != nil {
		err = fmt.Errorf("error with transaction: %w", err)
		return err
	}

	return tx.Commit(ctx)
}

func (d *Database) GetAlbums(ctx context.Context) ([]*models.Album, error) {
	res, err := d.Queries.GetAlbums(ctx)
	return wrapManyQueryError(res, err)
}

func (d *Database) GetAlbumsByID(ctx context.Context, id []string) ([]*models.Album, error) {
	res, err := d.Queries.GetAlbumsByID(ctx, id)
	return wrapManyQueryError(res, err)
}

func (d *Database) GetAlbumArtists(ctx context.Context, id string) ([]*models.Artist, error) {
	res, err := d.Queries.GetAlbumArtists(ctx, id)
	return wrapManyQueryError(res, err)
}

func (d *Database) GetAlbumImages(ctx context.Context, id string) ([]*models.Image, error) {
	res, err := d.Queries.GetAlbumImages(ctx, id)
	return wrapManyQueryError(res, err)
}

func (d *Database) GetAlbumTracks(ctx context.Context, id string) ([]*models.Track, error) {
	res, err := d.Queries.GetAlbumTracks(ctx, id)
	return wrapManyQueryError(res, err)
}

func (d *Database) CreateAlbum(ctx context.Context, album *database.CreateAlbumParams, imageURLs, artistIDs []string) (*models.Album, error) {
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

func (d *Database) GetArtistsByID(ctx context.Context, id []string) ([]*models.Artist, error) {
	res, err := d.Queries.GetArtistsByID(ctx, id)
	return wrapManyQueryError(res, err)
}

func (d *Database) GetArtistAlbums(ctx context.Context, id string) ([]*models.Album, error) {
	res, err := d.Queries.GetArtistAlbums(ctx, id)
	return wrapManyQueryError(res, err)
}

func (d *Database) GetArtistImages(ctx context.Context, id string) ([]*models.Image, error) {
	res, err := d.Queries.GetArtistImages(ctx, id)
	return wrapManyQueryError(res, err)
}

func (d *Database) GetArtistTracks(ctx context.Context, id string) ([]*models.Track, error) {
	res, err := d.Queries.GetArtistTracks(ctx, id)
	return wrapManyQueryError(res, err)
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

func (d *Database) GetImagesByURL(ctx context.Context, url []string) ([]*models.Image, error) {
	res, err := d.Queries.GetImagesByURL(ctx, url)
	return wrapManyQueryError(res, err)
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

func (d *Database) GetTracksByID(ctx context.Context, id []string) ([]*models.Track, error) {
	res, err := d.Queries.GetTracksByID(ctx, id)
	return wrapManyQueryError(res, err)
}

func (d *Database) GetTrackAlbum(ctx context.Context, id string) (*models.Album, error) {
	res, err := d.Queries.GetTrackAlbum(ctx, id)
	return wrapOneQueryError(res, err)
}

func (d *Database) GetTrackArtists(ctx context.Context, id string) ([]*models.Artist, error) {
	res, err := d.Queries.GetTrackArtists(ctx, id)
	return wrapManyQueryError(res, err)
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

func (d *Database) GetPlaylists(ctx context.Context) ([]*models.Playlist, error) {
	res, err := d.Queries.GetPlaylists(ctx)
	return wrapManyQueryError(res, err)
}

func (d *Database) GetPlaylistsByID(ctx context.Context, id []string) ([]*models.Playlist, error) {
	res, err := d.Queries.GetPlaylistsByID(ctx, id)
	return wrapManyQueryError(res, err)
}

func (d *Database) GetPlaylistImages(ctx context.Context, id string) ([]*models.Image, error) {
	res, err := d.Queries.GetPlaylistImages(ctx, id)
	return wrapManyQueryError(res, err)
}

func (d *Database) CreatePlaylist(ctx context.Context, playlist *database.CreatePlaylistParams, imageURLs []string) (*models.Playlist, error) {
	var p *models.Playlist

	err := d.withTX(ctx, func(q *database.Queries) error {
		var err error

		p, err = q.CreatePlaylist(ctx, playlist)
		if err != nil {
			return err
		}

		for _, url := range imageURLs {
			err = q.SetPlaylistImage(ctx, &database.SetPlaylistImageParams{
				PlaylistID: p.ID,
				ImageUrl:   url,
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

	return p, nil
}

func (d *Database) UpsertToken(ctx context.Context, token *database.UpsertTokenParams) (*models.Token, error) {
	t, err := d.Queries.UpsertToken(ctx, token)
	return wrapOneQueryError(t, err)
}

func (d *Database) GetToken(ctx context.Context) (*models.Token, error) {
	t, err := d.Queries.GetToken(ctx)
	return wrapOneQueryError(t, err)
}
