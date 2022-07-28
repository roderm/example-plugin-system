package graph

import (
	"context"
	"fmt"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/jmoiron/sqlx"
	"github.com/roderm/example-plugin-system/app/ent/entdb"
	"github.com/roderm/example-plugin-system/app/pkg/graph/v2/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
func GetSchema(ctx context.Context, db *sqlx.DB, extensions []string) graphql.ExecutableSchema {
	availableExtension := map[string]func(context.Context, *sqlx.DB) graphql.ExecutableSchema{}

	driver := entsql.OpenDB("sqlite3", db.DB)
	conn := entdb.NewClient(entdb.Driver(driver) /*prefix */)
	cfg := generated.Config{Resolvers: &Resolver{
		DB: conn,
	}}
	schema := generated.NewExecutableSchema(cfg)
	for _, name := range extensions {
		if ext, ok := availableExtension[name]; ok {
			if err := schema.Extend(ext(ctx, db)); err != nil {
				panic(err)
			}
		} else {
			fmt.Printf("Extension %s has not been loaded\n", name)
		}
	}
	return schema
}

type Resolver struct {
	DB *entdb.Client
}

func CheckPaging[K any](paging *K) *K {
	if paging == nil {
		return new(K)
	}
	return paging
}

func Options[R any]() []R {
	return []R{}
}
func AppendIfNotNil[R any](opts []R, option any, resolve func() R) []R {
	if option != nil {
		opts = append(opts, resolve())
	}
	return opts
}

func WithAfterFunc[O any](ctx context.Context, save func(context.Context) (O, error), cbs []func(O) error) (O, error) {
	res, err := save(ctx)
	if err != nil {
		return res, err
	}
	for _, cb := range cbs {
		if err := cb(res); err != nil {
			return res, err
		}
	}
	return res, nil
}
