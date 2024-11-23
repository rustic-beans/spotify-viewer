package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Track struct {
	ent.Schema
}

func (Track) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AtTimeMixin{},
	}
}

// Fields of the Track.
func (Track) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			NotEmpty().
			Unique().
			Immutable().
			Comment("The Spotify ID for the track"),
		field.Text("album_id").
			NotEmpty().
			Comment("The album on which the track appears"),
		field.JSON("available_markets", []string{}).
			Comment("A list of the countries in which the track can be played"),
		field.Int("disc_number").
			Optional().
			Comment("The disc number"),
		field.Int("duration_ms").
			Positive().
			Comment("The track length in milliseconds"),
		field.Bool("explicit").
			Default(false).
			Comment("Whether or not the track has explicit lyrics"),
		field.JSON("external_urls", &StringMap{}).
			Comment("Known external URLs for this track"),
		field.Text("href").
			NotEmpty().
			Comment("A link to the Web API endpoint providing full details of the track"),
		field.Bool("is_playable").
			Comment("If true, the track is playable in the given market"),
		field.Text("name").
			NotEmpty().
			Comment("The name of the track"),
		field.Int("popularity").
			Comment("The popularity of the track"),
		field.Text("preview_url").
			Optional().
			Comment("A link to a 30 second preview of the track"),
		field.Int("track_number").
			Comment("The number of the track"),
		field.Text("uri").
			NotEmpty().
			Comment("The Spotify URI for the track"),
	}
}

// Edges of the Track.
func (Track) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("artists", Artist.Type).
			Required().
			Ref("tracks"),
		edge.From("album", Album.Type).
			Required().
			Ref("tracks").
			Unique().
			Field("album_id"),
	}
}

func (Track) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}
