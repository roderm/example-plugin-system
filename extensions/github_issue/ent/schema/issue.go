package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	app_schema "github.com/roderm/example-plugin-system/app/ent/schema"
)

type Issue struct {
	ent.Schema
}

func (Issue) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("url"),
		field.Bool("is_pr").Default(false),
		field.String("status").Optional(),
	}
}
func (Issue) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("issue", app_schema.Todo.Type).Unique().StorageKey(edge.Column("id")),
	}
}
func (Issue) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tbl_issue"},
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
