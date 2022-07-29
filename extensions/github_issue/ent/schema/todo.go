package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
)

type Todo struct {
	ent.Schema
}

func (Todo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Skip(),
	}
}
