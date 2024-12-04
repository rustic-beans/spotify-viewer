// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/rustic-beans/spotify-viewer/ent/album"
	"github.com/rustic-beans/spotify-viewer/ent/schema"
)

// Album is the model entity for the Album schema.
type Album struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// The type of the album
	AlbumType album.AlbumType `json:"album_type,omitempty"`
	// The number of tracks in the album
	TotalTracks int `json:"total_tracks,omitempty"`
	// Known external URLs for this artist
	ExternalUrls *schema.StringMap `json:"external_urls,omitempty"`
	// A link to the Web API endpoint providing full details of the album
	Href string `json:"href,omitempty"`
	// The name of the album
	Name string `json:"name,omitempty"`
	// The date the album was first released
	ReleaseDate string `json:"release_date,omitempty"`
	// The precision with which release_date value is known
	ReleaseDatePrecision album.ReleaseDatePrecision `json:"release_date_precision,omitempty"`
	// The Spotify URI for the album
	URI string `json:"uri,omitempty"`
	// A list of the genres the album is associated with
	Genres []string `json:"genres,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AlbumQuery when eager-loading is set.
	Edges        AlbumEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AlbumEdges holds the relations/edges for other nodes in the graph.
type AlbumEdges struct {
	// Images holds the value of the images edge.
	Images []*Image `json:"images,omitempty"`
	// Artists holds the value of the artists edge.
	Artists []*Artist `json:"artists,omitempty"`
	// Tracks holds the value of the tracks edge.
	Tracks []*Track `json:"tracks,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedImages  map[string][]*Image
	namedArtists map[string][]*Artist
	namedTracks  map[string][]*Track
}

// ImagesOrErr returns the Images value or an error if the edge
// was not loaded in eager-loading.
func (e AlbumEdges) ImagesOrErr() ([]*Image, error) {
	if e.loadedTypes[0] {
		return e.Images, nil
	}
	return nil, &NotLoadedError{edge: "images"}
}

// ArtistsOrErr returns the Artists value or an error if the edge
// was not loaded in eager-loading.
func (e AlbumEdges) ArtistsOrErr() ([]*Artist, error) {
	if e.loadedTypes[1] {
		return e.Artists, nil
	}
	return nil, &NotLoadedError{edge: "artists"}
}

// TracksOrErr returns the Tracks value or an error if the edge
// was not loaded in eager-loading.
func (e AlbumEdges) TracksOrErr() ([]*Track, error) {
	if e.loadedTypes[2] {
		return e.Tracks, nil
	}
	return nil, &NotLoadedError{edge: "tracks"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Album) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case album.FieldExternalUrls, album.FieldGenres:
			values[i] = new([]byte)
		case album.FieldTotalTracks:
			values[i] = new(sql.NullInt64)
		case album.FieldID, album.FieldAlbumType, album.FieldHref, album.FieldName, album.FieldReleaseDate, album.FieldReleaseDatePrecision, album.FieldURI:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Album fields.
func (a *Album) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case album.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				a.ID = value.String
			}
		case album.FieldAlbumType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field album_type", values[i])
			} else if value.Valid {
				a.AlbumType = album.AlbumType(value.String)
			}
		case album.FieldTotalTracks:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_tracks", values[i])
			} else if value.Valid {
				a.TotalTracks = int(value.Int64)
			}
		case album.FieldExternalUrls:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field external_urls", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &a.ExternalUrls); err != nil {
					return fmt.Errorf("unmarshal field external_urls: %w", err)
				}
			}
		case album.FieldHref:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field href", values[i])
			} else if value.Valid {
				a.Href = value.String
			}
		case album.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case album.FieldReleaseDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field release_date", values[i])
			} else if value.Valid {
				a.ReleaseDate = value.String
			}
		case album.FieldReleaseDatePrecision:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field release_date_precision", values[i])
			} else if value.Valid {
				a.ReleaseDatePrecision = album.ReleaseDatePrecision(value.String)
			}
		case album.FieldURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uri", values[i])
			} else if value.Valid {
				a.URI = value.String
			}
		case album.FieldGenres:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field genres", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &a.Genres); err != nil {
					return fmt.Errorf("unmarshal field genres: %w", err)
				}
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Album.
// This includes values selected through modifiers, order, etc.
func (a *Album) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryImages queries the "images" edge of the Album entity.
func (a *Album) QueryImages() *ImageQuery {
	return NewAlbumClient(a.config).QueryImages(a)
}

// QueryArtists queries the "artists" edge of the Album entity.
func (a *Album) QueryArtists() *ArtistQuery {
	return NewAlbumClient(a.config).QueryArtists(a)
}

// QueryTracks queries the "tracks" edge of the Album entity.
func (a *Album) QueryTracks() *TrackQuery {
	return NewAlbumClient(a.config).QueryTracks(a)
}

// Update returns a builder for updating this Album.
// Note that you need to call Album.Unwrap() before calling this method if this Album
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Album) Update() *AlbumUpdateOne {
	return NewAlbumClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Album entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Album) Unwrap() *Album {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Album is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Album) String() string {
	var builder strings.Builder
	builder.WriteString("Album(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("album_type=")
	builder.WriteString(fmt.Sprintf("%v", a.AlbumType))
	builder.WriteString(", ")
	builder.WriteString("total_tracks=")
	builder.WriteString(fmt.Sprintf("%v", a.TotalTracks))
	builder.WriteString(", ")
	builder.WriteString("external_urls=")
	builder.WriteString(fmt.Sprintf("%v", a.ExternalUrls))
	builder.WriteString(", ")
	builder.WriteString("href=")
	builder.WriteString(a.Href)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(a.Name)
	builder.WriteString(", ")
	builder.WriteString("release_date=")
	builder.WriteString(a.ReleaseDate)
	builder.WriteString(", ")
	builder.WriteString("release_date_precision=")
	builder.WriteString(fmt.Sprintf("%v", a.ReleaseDatePrecision))
	builder.WriteString(", ")
	builder.WriteString("uri=")
	builder.WriteString(a.URI)
	builder.WriteString(", ")
	builder.WriteString("genres=")
	builder.WriteString(fmt.Sprintf("%v", a.Genres))
	builder.WriteByte(')')
	return builder.String()
}

// NamedImages returns the Images named value or an error if the edge was not
// loaded in eager-loading with this name.
func (a *Album) NamedImages(name string) ([]*Image, error) {
	if a.Edges.namedImages == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := a.Edges.namedImages[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (a *Album) appendNamedImages(name string, edges ...*Image) {
	if a.Edges.namedImages == nil {
		a.Edges.namedImages = make(map[string][]*Image)
	}
	if len(edges) == 0 {
		a.Edges.namedImages[name] = []*Image{}
	} else {
		a.Edges.namedImages[name] = append(a.Edges.namedImages[name], edges...)
	}
}

// NamedArtists returns the Artists named value or an error if the edge was not
// loaded in eager-loading with this name.
func (a *Album) NamedArtists(name string) ([]*Artist, error) {
	if a.Edges.namedArtists == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := a.Edges.namedArtists[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (a *Album) appendNamedArtists(name string, edges ...*Artist) {
	if a.Edges.namedArtists == nil {
		a.Edges.namedArtists = make(map[string][]*Artist)
	}
	if len(edges) == 0 {
		a.Edges.namedArtists[name] = []*Artist{}
	} else {
		a.Edges.namedArtists[name] = append(a.Edges.namedArtists[name], edges...)
	}
}

// NamedTracks returns the Tracks named value or an error if the edge was not
// loaded in eager-loading with this name.
func (a *Album) NamedTracks(name string) ([]*Track, error) {
	if a.Edges.namedTracks == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := a.Edges.namedTracks[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (a *Album) appendNamedTracks(name string, edges ...*Track) {
	if a.Edges.namedTracks == nil {
		a.Edges.namedTracks = make(map[string][]*Track)
	}
	if len(edges) == 0 {
		a.Edges.namedTracks[name] = []*Track{}
	} else {
		a.Edges.namedTracks[name] = append(a.Edges.namedTracks[name], edges...)
	}
}

// Albums is a parsable slice of Album.
type Albums []*Album
