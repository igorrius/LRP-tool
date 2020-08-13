package parser

import (
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
		Action: func(ctx *cli.Context) error {
			log := logger.FromContextOrNew(ctx.Context)
			conf := project.NewConfig("./config/project", log)

			return parser.NewParser(
				ctx.Context,
				log,
				conf.Configurations().Normalized(),
			).Run(
				//Parser options
				parser.Async(false),
				parser.Storage(storage.NewStorage(ctx.Context)),
				parser.CacheDir("./.cache"),
			)
		},
	}
}
