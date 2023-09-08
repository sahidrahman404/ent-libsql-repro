package schema

import (
	"todo/ent/schema/pksuid"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Exercise holds the schema definition for the Exercise entity.
type Exercise struct {
	ent.Schema
}

// Fields of the Exercise.
func (Exercise) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("image").Optional().Nillable(),
		field.String("how_to").Optional().Nillable(),
	}
}

// Edges of the Exercise.
func (Exercise) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("muscles_groups", MusclesGroup.Type).
			Annotations(entsql.OnDelete(entsql.Cascade), entgql.RelayConnection()),
	}
}

func (Exercise) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (Exercise) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pksuid.MixinWithPrefix("EX"),
	}
}
