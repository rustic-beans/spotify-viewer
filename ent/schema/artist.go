package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
		field.JSON("genres", []string{}).
			Comment("A list of genres the artist is associated with.  For example, \"Prog Rock\" or \"Post-Grunge\".  If not yet classified, the slice is empty."),
	}
}

// Edges of the Artist.
func (Artist) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("albums", Album.Type),
		edge.To("tracks", Track.Type),
		edge.To("images", Image.Type).Required(),
	}
}

func (Artist) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}
