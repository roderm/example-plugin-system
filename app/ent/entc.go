//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	templates := []*gen.Template{
		entgql.CollectionTemplate,
		entgql.EnumTemplate,
		entgql.NodeTemplate,
		entgql.PaginationTemplate,
		entgql.TransactionTemplate,
		entgql.EdgeTemplate,
		// override entgql.MutationInputTemplate
		MutationInputTemplate(),
	}
	extGql, err := entgql.NewExtension(
		entgql.WithConfigPath("./gqlgen.yml"),
		entgql.WithTemplates(templates...),
		entgql.WithSchemaPath("./api/ent.graphql"),
		entgql.WithWhereFilters(true),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	// rt := reflect.TypeOf(uuid.UUID{})

	if err := entc.Generate("./ent/schema", &gen.Config{
		Target:  "./ent/entdb/",
		Package: "github.com/roderm/example-plugin-system/app/ent/entdb",
		// IDType: &field.TypeInfo{
		// 	Type:    field.TypeUUID,
		// 	Ident:   rt.String(),
		// 	PkgPath: rt.PkgPath(),
		// },
	},
		entc.Extensions(extGql),
	// entc.TemplateDir("./ent/template"),
	); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
func MutationInputTemplate() *gen.Template {
	return gen.MustParse(gen.NewTemplate("./ent/template/mutation_input.tmpl").
		Funcs(entgql.TemplateFuncs).ParseFiles("./ent/template/mutation_input.tmpl"))
}
