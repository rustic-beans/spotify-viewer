package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Image holds the schema definition for the Image entity.
type Image struct {
	ent.Schema
}

// Fields of the Todo.
func (Image) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			NotEmpty().
			Unique(),
		field.String("url").NotEmpty(),
		field.Int("width").Positive(),
		field.Int("height").Positive(),
	}
}

// Edges of the Todo.
func (Image) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("albums", Album.Type).Ref("images"),
		edge.From("artists", Artist.Type).Ref("images"),
	}
}

func (Image) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}
