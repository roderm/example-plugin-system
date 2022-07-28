package main

//go:generate go run ent/entc.go
//go:generate go run github.com/99designs/gqlgen generate --config gqlgen.yml

import (
	_ "entgo.io/contrib/entgql"
	_ "entgo.io/ent"
	_ "github.com/99designs/gqlgen"
)
