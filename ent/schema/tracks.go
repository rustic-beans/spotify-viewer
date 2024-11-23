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

func (Track) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			NotEmpty().
			Unique().
			Immutable(),
		field.String("name"),
		field.Strings("artists"),
		field.Strings("artists_genres"),
		field.String("album_name"),
		field.String("album_image_uri"),
		field.Int("duration_ms"),
		field.String("uri"),
	}
}

func (Track) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("albums", Album.Type).
			Ref("tracks").
			Unique(),
	}
}

func (Track) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}
