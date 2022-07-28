package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("firstname").Optional(),
		field.String("lastname").Optional(),
		field.String("email"),
	}
}
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("todos", Todo.Type).StorageKey(edge.Column("user_id")),
	}
}
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tbl_user"},
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
