package services

import (
	"context"
	"fmt"

	"github.com/cockroachdb/errors"
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
	GetAlbumTracks(ctx context.Context, id string) ([]*models.Track, error)
	CreateAlbum(ctx context.Context, album *database.CreateAlbumParams, artistIDs []string) (*models.Album, error)

	GetArtists(ctx context.Context) ([]*models.Artist, error)
	GetArtistsByID(ctx context.Context, id []string) ([]*models.Artist, error)
	GetArtistAlbums(ctx context.Context, id string) ([]*models.Album, error)
	GetArtistTracks(ctx context.Context, id string) ([]*models.Track, error)
	CreateArtist(ctx context.Context, artist *database.CreateArtistParams) (*models.Artist, error)

	GetTracks(ctx context.Context) ([]*models.Track, error)
	GetTracksByID(ctx context.Context, id []string) ([]*models.Track, error)
	GetTrackAlbum(ctx context.Context, id string) (*models.Album, error)
	GetTrackArtists(ctx context.Context, id string) ([]*models.Artist, error)
	CreateTrack(ctx context.Context, track *database.CreateTrackParams, artistIDs []string) (*models.Track, error)

	GetPlaylists(ctx context.Context) ([]*models.Playlist, error)
	GetPlaylistsByID(ctx context.Context, id []string) ([]*models.Playlist, error)
	CreatePlaylist(ctx context.Context, playlist *database.CreatePlaylistParams) (*models.Playlist, error)

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

func wrapOneQueryError[T any](result *T, err error, wrapMessage string) (*T, error) {
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	return result, errors.Wrap(err, wrapMessage)
}

func wrapManyQueryError[T any](result []*T, err error, wrapMessage string) ([]*T, error) {
	if err == pgx.ErrNoRows {
		return nil, nil
	}

	return result, errors.Wrap(err, wrapMessage)
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

	qtx := d.WithTx(tx)
	if err = fn(qtx); err != nil {
		return errors.Wrap(err, "error while executing transaction")
	}

	return tx.Commit(ctx)
}

func (d *Database) GetAlbums(ctx context.Context) ([]*models.Album, error) {
	res, err := d.Queries.GetAlbums(ctx)
	return wrapManyQueryError(res, err, "error getting albums from database")
}

func (d *Database) GetAlbumsByID(ctx context.Context, id []string) ([]*models.Album, error) {
	res, err := d.Queries.GetAlbumsByID(ctx, id)
	return wrapManyQueryError(res, err, fmt.Sprintf("error getting albums by ids %v from database", id))
}

func (d *Database) GetAlbumArtists(ctx context.Context, id string) ([]*models.Artist, error) {
	res, err := d.Queries.GetAlbumArtists(ctx, id)
	return wrapManyQueryError(res, err, fmt.Sprintf("error getting album artists by id %s from database", id))
}

func (d *Database) GetAlbumTracks(ctx context.Context, id string) ([]*models.Track, error) {
	res, err := d.Queries.GetAlbumTracks(ctx, id)
	return wrapManyQueryError(res, err, fmt.Sprintf("error getting album tracks by id %s from database", id))
}

func (d *Database) CreateAlbum(ctx context.Context, album *database.CreateAlbumParams, artistIDs []string) (a *models.Album, err error) {
	err = d.withTX(ctx, func(q *database.Queries) error {
		a, err = q.CreateAlbum(ctx, album)
		if err != nil {
			return errors.Wrap(err, "failed to run create album query")
		}

		for _, id := range artistIDs {
			err = q.SetArtistAlbum(ctx, &database.SetArtistAlbumParams{
				AlbumID:  a.ID,
				ArtistID: id,
			})
			if err != nil {
				return errors.Wrap(err, "failed to run set artist album query")
			}
		}

		return nil
	})

	return a, errors.Wrap(err, "error creating album")
}

func (d *Database) GetArtists(ctx context.Context) ([]*models.Artist, error) {
	res, err := d.Queries.GetArtists(ctx)
	return wrapManyQueryError(res, err, "error getting artists from database")
}

func (d *Database) GetArtistsByID(ctx context.Context, id []string) ([]*models.Artist, error) {
	res, err := d.Queries.GetArtistsByID(ctx, id)
	return wrapManyQueryError(res, err, fmt.Sprintf("error getting artists by ids %v from database", id))
}

func (d *Database) GetArtistAlbums(ctx context.Context, id string) ([]*models.Album, error) {
	res, err := d.Queries.GetArtistAlbums(ctx, id)
	return wrapManyQueryError(res, err, fmt.Sprintf("error getting artist albums by id %s from database", id))
}

func (d *Database) GetArtistTracks(ctx context.Context, id string) ([]*models.Track, error) {
	res, err := d.Queries.GetArtistTracks(ctx, id)
	return wrapManyQueryError(res, err, fmt.Sprintf("error getting artist tracks by id %s from database", id))
}

func (d *Database) CreateArtist(ctx context.Context, artist *database.CreateArtistParams) (a *models.Artist, err error) {
	err = d.withTX(ctx, func(q *database.Queries) error {
		a, err = q.CreateArtist(ctx, artist)
		if err != nil {
			return errors.Wrap(err, "failed to run create artist query")
		}

		return nil
	})

	return a, errors.Wrap(err, "error creating artist")
}

func (d *Database) GetTracks(ctx context.Context) ([]*models.Track, error) {
	res, err := d.Queries.GetTracks(ctx)
	return wrapManyQueryError(res, err, "error getting tracks from database")
}

func (d *Database) GetTracksByID(ctx context.Context, id []string) ([]*models.Track, error) {
	res, err := d.Queries.GetTracksByID(ctx, id)
	return wrapManyQueryError(res, err, fmt.Sprintf("error getting tracks by ids %v from database", id))
}

func (d *Database) GetTrackAlbum(ctx context.Context, id string) (*models.Album, error) {
	res, err := d.Queries.GetTrackAlbum(ctx, id)
	return wrapOneQueryError(res, err, fmt.Sprintf("error getting track album by id %s from database", id))
}

func (d *Database) GetTrackArtists(ctx context.Context, id string) ([]*models.Artist, error) {
	res, err := d.Queries.GetTrackArtists(ctx, id)
	return wrapManyQueryError(res, err, fmt.Sprintf("error getting track artists by id %s from database", id))
}

func (d *Database) CreateTrack(ctx context.Context, track *database.CreateTrackParams, artistIDs []string) (t *models.Track, err error) {
	err = d.withTX(ctx, func(q *database.Queries) error {
		t, err = q.CreateTrack(ctx, track)
		if err != nil {
			return errors.Wrap(err, "failed to run create track query")
		}

		for _, id := range artistIDs {
			err = q.SetArtistTrack(ctx, &database.SetArtistTrackParams{
				ArtistID: id,
				TrackID:  t.ID,
			})
			if err != nil {
				return errors.Wrap(err, "failed to run set artist track query")
			}
		}

		return nil
	})

	return t, errors.Wrap(err, "error creating track")
}

func (d *Database) GetPlaylists(ctx context.Context) ([]*models.Playlist, error) {
	res, err := d.Queries.GetPlaylists(ctx)
	return wrapManyQueryError(res, err, "error getting playlists from database")
}

func (d *Database) GetPlaylistsByID(ctx context.Context, id []string) ([]*models.Playlist, error) {
	res, err := d.Queries.GetPlaylistsByID(ctx, id)
	return wrapManyQueryError(res, err, fmt.Sprintf("error getting playlists by ids %v from database", id))
}

func (d *Database) CreatePlaylist(ctx context.Context, playlist *database.CreatePlaylistParams) (p *models.Playlist, err error) {
	err = d.withTX(ctx, func(q *database.Queries) error {
		p, err = q.CreatePlaylist(ctx, playlist)
		if err != nil {
			return errors.Wrap(err, "failed to run create playlist query")
		}

		return nil
	})

	return p, errors.Wrap(err, "error creating playlist")
}

func (d *Database) UpsertToken(ctx context.Context, token *database.UpsertTokenParams) (*models.Token, error) {
	t, err := d.Queries.UpsertToken(ctx, token)
	return wrapOneQueryError(t, err, "error upserting token")
}

func (d *Database) GetToken(ctx context.Context) (*models.Token, error) {
	t, err := d.Queries.GetToken(ctx)
	return wrapOneQueryError(t, err, "error getting token")
}
