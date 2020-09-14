package parser

import (
	"github.com/choizydev/LRP-tool/internal/cmd/flag"
	"github.com/choizydev/LRP-tool/internal/logger"
	"github.com/choizydev/LRP-tool/internal/parser"
	"github.com/choizydev/LRP-tool/internal/project"
	"github.com/choizydev/LRP-tool/internal/storage"
	"github.com/urfave/cli/v2"
)

func Run() *cli.Command {
	return &cli.Command{
		Name:  "run",
		Usage: "Parse data from remote sources using projects configuration",
		Flags: []cli.Flag{
			flag.PostgreSQL(),
			flag.DummyStorage(),
		},
		Action: func(ctx *cli.Context) error {
			log := logger.FromContextOrNew(ctx.Context)
			conf := project.NewConfig("./config/project", log)

			return parser.NewParser(
				log,
				conf.Configurations().Normalized(),
			).Run(
				//Parser options
				parser.Async(false),
				buildStorageOption(ctx),
				parser.CacheDir("./.cache"),
			)
		},
	}
}

//for now the application can use only an one storage
func buildStorageOption(ctx *cli.Context) parser.Option {
	storages := make([]storage.Option, 0, 8)

	// enable dummy storage if an appropriate flag is set
	if ctx.Bool(flag.DummyStorageName) {
		storages = append(storages, storage.WithEngine(
			storage.NewDummyEngine(ctx.Context),
		))
	}

	// enable PostgreSQL storage if an appropriate flag is set
	if dsn := ctx.String(flag.PostgreSQLStorageName); dsn != "" {
		storages = append(storages, storage.WithEngine(
			storage.NewPostgresqlEngine(ctx.Context, dsn),
		))
	}

	return parser.Storage(storage.NewStorage(
		ctx.Context,
		storages...,
	))
}
