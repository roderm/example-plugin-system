package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/roderm/example-plugin-system/app/pkg/graph/playground/apollo"
	"github.com/roderm/example-plugin-system/app/pkg/graph/v2"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: "config"},
		altsrc.NewIntFlag(&cli.IntFlag{
			Name:  "port",
			Value: 3080,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  "sqlite",
			Value: "example.db",
		}),
		altsrc.NewStringSliceFlag(&cli.StringSliceFlag{
			Name: "ext",
		}),
	}

	app.Action = func(ctx *cli.Context) error {
		// setup db
		db, err := sqlx.Open("sqlite3", fmt.Sprintf("file:db//%s?mode=memory&cache=shared&_fk=1", ctx.String("sqlite")))
		if err != nil {
			return err
		}
		schema := graph.GetSchema(ctx.Context, db, ctx.StringSlice("ext"))
		address := fmt.Sprintf("0.0.0.0:%d", ctx.Int("port"))
		h := http.NewServeMux()
		h.Handle("/api", apollo.Handler("GraphQL playground", "/"))
		h.Handle("/", graph.Wrap(schema))
		return http.ListenAndServe(address, h)
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
