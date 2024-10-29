// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/rustic-beans/spotify-viewer/ent/schema/pulid"
	"github.com/rustic-beans/spotify-viewer/ent/track"
)

// Track is the model entity for the Track schema.
type Track struct {
	config `json:"-"`
	// ID of the ent.
	ID pulid.ID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// TrackID holds the value of the "track_id" field.
	TrackID pulid.ID `json:"track_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Artists holds the value of the "artists" field.
	Artists []string `json:"artists,omitempty"`
	// ArtistsGenres holds the value of the "artists_genres" field.
	ArtistsGenres []string `json:"artists_genres,omitempty"`
	// AlbumName holds the value of the "album_name" field.
	AlbumName string `json:"album_name,omitempty"`
	// AlbumImageURI holds the value of the "album_image_uri" field.
	AlbumImageURI string `json:"album_image_uri,omitempty"`
	// DurationMs holds the value of the "duration_ms" field.
	DurationMs int `json:"duration_ms,omitempty"`
	// URI holds the value of the "uri" field.
	URI          string `json:"uri,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Track) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case track.FieldArtists, track.FieldArtistsGenres:
			values[i] = new([]byte)
		case track.FieldID, track.FieldTrackID:
			values[i] = new(pulid.ID)
		case track.FieldDurationMs:
			values[i] = new(sql.NullInt64)
		case track.FieldName, track.FieldAlbumName, track.FieldAlbumImageURI, track.FieldURI:
			values[i] = new(sql.NullString)
		case track.FieldCreatedAt, track.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Track fields.
func (t *Track) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case track.FieldID:
			if value, ok := values[i].(*pulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				t.ID = *value
			}
		case track.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case track.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		case track.FieldTrackID:
			if value, ok := values[i].(*pulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field track_id", values[i])
			} else if value != nil {
				t.TrackID = *value
			}
		case track.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case track.FieldArtists:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field artists", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &t.Artists); err != nil {
					return fmt.Errorf("unmarshal field artists: %w", err)
				}
			}
		case track.FieldArtistsGenres:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field artists_genres", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &t.ArtistsGenres); err != nil {
					return fmt.Errorf("unmarshal field artists_genres: %w", err)
				}
			}
		case track.FieldAlbumName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field album_name", values[i])
			} else if value.Valid {
				t.AlbumName = value.String
			}
		case track.FieldAlbumImageURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field album_image_uri", values[i])
			} else if value.Valid {
				t.AlbumImageURI = value.String
			}
		case track.FieldDurationMs:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration_ms", values[i])
			} else if value.Valid {
				t.DurationMs = int(value.Int64)
			}
		case track.FieldURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uri", values[i])
			} else if value.Valid {
				t.URI = value.String
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Track.
// This includes values selected through modifiers, order, etc.
func (t *Track) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// Update returns a builder for updating this Track.
// Note that you need to call Track.Unwrap() before calling this method if this Track
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Track) Update() *TrackUpdateOne {
	return NewTrackClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Track entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Track) Unwrap() *Track {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Track is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Track) String() string {
	var builder strings.Builder
	builder.WriteString("Track(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("track_id=")
	builder.WriteString(fmt.Sprintf("%v", t.TrackID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteString(", ")
	builder.WriteString("artists=")
	builder.WriteString(fmt.Sprintf("%v", t.Artists))
	builder.WriteString(", ")
	builder.WriteString("artists_genres=")
	builder.WriteString(fmt.Sprintf("%v", t.ArtistsGenres))
	builder.WriteString(", ")
	builder.WriteString("album_name=")
	builder.WriteString(t.AlbumName)
	builder.WriteString(", ")
	builder.WriteString("album_image_uri=")
	builder.WriteString(t.AlbumImageURI)
	builder.WriteString(", ")
	builder.WriteString("duration_ms=")
	builder.WriteString(fmt.Sprintf("%v", t.DurationMs))
	builder.WriteString(", ")
	builder.WriteString("uri=")
	builder.WriteString(t.URI)
	builder.WriteByte(')')
	return builder.String()
}

// Tracks is a parsable slice of Track.
type Tracks []*Track
