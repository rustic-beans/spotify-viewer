package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Artist holds the schema definition for the Artist entity.
type Artist struct {
	ent.Schema
}

// Fields of the Artist.
func (Artist) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			NotEmpty().
			Unique().
			Immutable().
			Comment("The Spotify ID for the artist"),
		field.JSON("external_urls", &StringMap{}).
			Comment("Known external URLs for this artist"),
		field.Text("href").
			NotEmpty().
			Comment("A link to the Web API endpoint providing full details of the artist"),
		field.Text("name").
			NotEmpty().
			Comment("The name of the artist"),
		field.Text("uri").
			NotEmpty().
			Comment("The Spotify URI for the artist"),
	}
}

// Edges of the Artist.
func (Artist) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("albums", Album.Type),
		edge.To("tracks", Track.Type),
	}
}
