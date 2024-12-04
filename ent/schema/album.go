package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Album holds the schema definition for the Album entity.
type Album struct {
	ent.Schema
}

// Fields of the Album.
func (Album) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			NotEmpty().
			Unique().
			Immutable(),
		field.Enum("album_type").
			Values("album", "single", "compilation").
			Comment("The type of the album"),
		field.Int("total_tracks").
			Comment("The number of tracks in the album"),
		field.JSON("external_urls", &StringMap{}).
			Comment("Known external URLs for this artist"),
		field.Text("href").
			NotEmpty().
			Comment("A link to the Web API endpoint providing full details of the album"),
		field.Text("name").
			NotEmpty().
			Comment("The name of the album"),
		field.Text("release_date").
			NotEmpty().
			Comment("The date the album was first released"),
		field.Enum("release_date_precision").
			Values("year", "month", "day").
			Comment("The precision with which release_date value is known"),
		field.Text("uri").
			NotEmpty().
			Comment("The Spotify URI for the album"),
		field.JSON("genres", []string{}).
			Comment("A list of the genres the album is associated with"),
	}
}

// Edges of the Album.
func (Album) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("images", Image.Type).Required(),
		edge.From("artists", Artist.Type).Ref("albums"),
		edge.To("tracks", Track.Type),
	}
}

func (Album) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}