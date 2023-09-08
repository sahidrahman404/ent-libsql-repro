package schema

import (
	"todo/ent/schema/pksuid"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MusclesGroup holds the schema definition for the MusclesGroup entity.
type MusclesGroup struct {
	ent.Schema
}

// Fields of the MusclesGroup.
func (MusclesGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("image"),
	}
}

// Edges of the MusclesGroup.
func (MusclesGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("exercises", Exercise.Type).
			Ref("muscles_groups").
			Annotations(entgql.RelayConnection()),
	}
}

func (MusclesGroup) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (MusclesGroup) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pksuid.MixinWithPrefix("MG"),
	}
}
