package services

import (
	"context"

	"github.com/rustic-beans/spotify-viewer/ent"
	"github.com/rustic-beans/spotify-viewer/internal/models"
)

type Database struct {
	client *ent.Client
}

func NewDatabase(client *ent.Client) *Database {
	return &Database{client: client}
}

func (d *Database) GetArtist(ctx context.Context, id string) (*models.Artist, error) {
	artist, err := d.client.Artist.Get(ctx, id)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return artist, nil
}

func (d *Database) GetAlbum(ctx context.Context, id string) (*models.Album, error) {
	album, err := d.client.Album.Get(ctx, id)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return album, nil
}

func (d *Database) GetTrack(ctx context.Context, id string) (*models.Track, error) {
	track, err := d.client.Track.Get(ctx, id)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	return track, nil
}

func (d *Database) SaveArtist(ctx context.Context, artist *models.Artist, images []*models.Image) (*models.Artist, error) {
	if err := d.UpsertImages(ctx, images); err != nil {
		return nil, err
	}

	artistModel, err := d.client.Artist.Create().
		SetID(artist.ID).
		SetExternalUrls(artist.ExternalUrls).
		SetHref(artist.Href).
		SetName(artist.Name).
		SetURI(artist.URI).
		SetGenres(artist.Genres).
		AddImages(images...).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return artistModel, nil
}

func (d *Database) SaveAlbum(ctx context.Context, album *models.Album, images []*models.Image, artistIDs []string) (*models.Album, error) {
	if err := d.UpsertImages(ctx, images); err != nil {
		return nil, err
	}

	albumModel, err := d.client.Album.Create().
		SetID(album.ID).
		SetAlbumType(album.AlbumType).
		SetTotalTracks(album.TotalTracks).
		SetExternalUrls(album.ExternalUrls).
		SetHref(album.Href).
		SetName(album.Name).
		SetReleaseDate(album.ReleaseDate).
		SetReleaseDatePrecision(album.ReleaseDatePrecision).
		SetURI(album.URI).
		SetGenres(album.Genres).
		AddImages(images...).
		AddArtistIDs(artistIDs...).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return albumModel, nil
}

func (d *Database) UpsertImages(ctx context.Context, images []*models.Image) error {
	builders := make([]*ent.ImageCreate, len(images))

	for i, image := range images {
		builders[i] = d.client.Image.Create().
			SetID(image.ID).
			SetHeight(image.Height).
			SetWidth(image.Width).
			SetURL(image.URL)
	}

	err := d.client.Image.
		CreateBulk(builders...).
		OnConflict().
		UpdateNewValues().
		Exec(ctx)

	return err
}

func (d *Database) SaveTrack(ctx context.Context, track *models.Track, artistIDs []string) (*models.Track, error) {
	trackModel, err := d.client.Track.Create().
		SetID(track.ID).
		SetAlbumID(track.AlbumID).
		SetDiscNumber(track.DiscNumber).
		SetDurationMs(track.DurationMs).
		SetExternalUrls(track.ExternalUrls).
		SetHref(track.Href).
		SetName(track.Name).
		SetPopularity(track.Popularity).
		SetPreviewURL(track.PreviewURL).
		SetTrackNumber(track.TrackNumber).
		SetURI(track.URI).
		AddArtistIDs(artistIDs...).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return trackModel, nil
}

func (d *Database) GetTracks(ctx context.Context) ([]*models.Track, error) {
	tracks, err := d.client.Track.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return tracks, nil
}

func (d *Database) GetAlbums(ctx context.Context) ([]*models.Album, error) {
	albums, err := d.client.Album.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return albums, nil
}

func (d *Database) GetArtists(ctx context.Context) ([]*models.Artist, error) {
	artists, err := d.client.Artist.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func (d *Database) GetImages(ctx context.Context) ([]*models.Image, error) {
	images, err := d.client.Image.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return images, nil
}
