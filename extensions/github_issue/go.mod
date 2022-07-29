module github.com/roderm/example-plugin-system/extensions/github_issue

go 1.18

replace github.com/roderm/example-plugin-system v0.0.0 => ../../

require (
	entgo.io/contrib v0.2.1-0.20220513120443-ee1f1c4f1d3b
	entgo.io/ent v0.10.2-0.20220512043615-f2e0bef7a803
	github.com/99designs/gqlgen v0.17.13
)
