package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Todo struct {
	ent.Schema
}

func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("description"),
		field.Time("done").Optional(),
	}
}
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("todos").Required().Unique(),
	}
}
func (Todo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tbl_odo"},
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
